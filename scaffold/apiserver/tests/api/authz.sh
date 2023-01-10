#!/usr/bin/env bash

INSECURE_SERVER="127.0.0.1:8080"
#INSECURE_SERVER="127.0.0.1:30080"  # Docker port
SECURE_SERVER="127.0.0.1:8443"
INSECURE_AUTHZ_SERVER="127.0.0.1:8090"
#INSECURE_AUTHZ_SERVER="127.0.0.1:30090"

Header="-HContent-Type: application/json"
CCURL="curl -f -s -XPOST" # Create
UCURL="curl -f -s -XPUT" # Update
RCURL="curl -f -s -XGET" # Get
DCURL="curl -f -s -XDELETE" # Delete

insecure::login()
{
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/login \
    -d'{"username":"admin","password":"Admin@2021"}' |  jq '.token' | sed 's/"//g'
}

insecure::authz-noauth()
{
  # 1. admin login
  token="-HAuthorization: Bearer $(insecure::login)"

  # 2 如果有 wkpolicy 策略先清空
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/policies/wkpolicy; echo

  # 3 创建 wkpolicy 策略
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"wkpolicy"},"policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 4. 调用 /v1/authz 完成资源鉴权
  $CCURL "${Header}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:maria","action":"delete","resource":"resources:articles:ladon-introduction","context":{"remoteIPAddress":"192.168.0.5"}}'
}

insecure::authz-auth()
{
  # 1. admin login
  token="-HAuthorization: Bearer $(insecure::login)"

  # 2 如果有 wkpolicy 策略先清空
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/policies/wkpolicy; echo

  # 3 创建 wkpolicy 策略
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"wkpolicy"},"policy":{"description":"One policy to rule them all.","subjects":["users:<peter|ken>","users:maria","groups:admins"],"actions":["delete","<create|update>"],"effect":"allow","resources":["resources:articles:<.*>","resources:printer"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo
  
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题

  # 2. 用 admin，如果有 test00、test01 用户先清空
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test00; echo
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/test01; echo

  # 3. 用 admin，创建 test00、test01 用户
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"test01"},"password":"User@2022","nickname":"01","email":"test01@gmail.com","phone":"1306280xxxx"}'; echo

  # 4. 创建 test00 的 basic auth header
  basic00=$(echo -n 'test00:User@2022'|base64)
  Header00="-HAuthorization: Basic ${basic00}"

  # 5. 用 test00，列出所有用户
  ${RCURL} "${Header00}" "http://${INSECURE_SERVER}/v1/users?offset=0&limit=10"; echo
  
  # 4. 调用 /v1/authz 完成资源鉴权
  $CCURL "${Header}" "${Header00}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:maria","action":"delete","resource":"resources:articles:ladon-introduction","context":{"remoteIPAddress":"192.168.0.5"}}'

}

insecure::question()
{
  ## create users
  # 1. 创建 admin 的 basic auth header, admin 密码为 "Admin@2021"
  basic=$(echo -n 'admin:Admin@2021'|base64)
  HeaderBasic="-HAuthorization: Basic ${basic}"  # 注意 -H 与 Authorization 间不能有空格，否则解析会有问题
  
  # 2. 用 admin，如果有 student00、student01 用户先清空
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/student00; echo
  ${DCURL} "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users/student01; echo

  # 3. 用 admin，创建 test00、test01 用户
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"student00"},"password":"User@2022","nickname":"00","email":"test00@gmail.com","phone":"1306280xxxx"}'; echo
  ${CCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_SERVER}/v1/users \
    -d'{"metadata":{"name":"student01"},"password":"User@2022","nickname":"01","email":"test01@gmail.com","phone":"1306280xxxx"}'; echo
  
  # 4. 创建 test00 的 basic auth header
  basic00=$(echo -n 'student00:User@2022'|base64)
  Header00="-HAuthorization: Basic ${basic00}"

  
  ## create policy
  # 1. admin login
  token="-HAuthorization: Bearer $(insecure::login)"

  # 2 如果有 wkpolicy 策略先清空
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/policies/wkpolicy; echo
  ${DCURL} "${token}" http://${INSECURE_SERVER}/v1/policies/wkpolicy2; echo

  # 3 创建 wkpolicy 策略
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"wkpolicy"},"policy":{"description":"first policy.","subjects":["users:<student.*>","users:admin"],"actions":["<list|get>","<create|update>"],"effect":"allow","resources":["resources:questions:<.*>"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 4 创建 wkpolicy2 策略
  ${CCURL} "${Header}" "${token}" http://${INSECURE_SERVER}/v1/policies \
    -d'{"metadata":{"name":"wkpolicy2"},"policy":{"description":"second policy.","subjects":["users:admin"],"actions":["delete"],"effect":"allow","resources":["resources:questions:<.*>"],"conditions":{"remoteIPAddress":{"type":"CIDRCondition","options":{"cidr":"192.168.0.1/16"}}}}}'; echo

  # 5. use student00 to log in /v1/authz 完成资源鉴权 and create question
  $CCURL "${Header}" "${Header00}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:student","action":"create","resource":"resources:questions:q1","context":{"remoteIPAddress":"192.168.0.5"}, "question":{"metadata":{"name":"q1"},"studentname":"student01","content":"question1"}}'; echo
  # 6. use student00 to log in /v1/authz and get the question list
  $RCURL "${Header}" "${Header00}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:student","action":"list","resource":"resources:questions:q1","context":{"remoteIPAddress":"192.168.0.5"}}'; echo

  # 7. use admin to delete question
  ${DCURL} "${Header}" "${HeaderBasic}" http://${INSECURE_AUTHZ_SERVER}/v1/authz/q1 \
    -d'{"subject":"users:admin","action":"list","resource":"resources:questions:q1","context":{"remoteIPAddress":"192.168.0.5"}}'; echo

  # 8. use admin to get the question list
  $RCURL "${HeaderBasic}" http://${INSECURE_AUTHZ_SERVER}/v1/authz \
    -d'{"subject":"users:admin","action":"list","resource":"resources:questions:q1","context":{"remoteIPAddress":"192.168.0.5"}}'; echo

    
}

if [[ "$*" =~ insecure:: ]];then
  eval $*
fi