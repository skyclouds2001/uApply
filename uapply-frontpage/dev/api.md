
 <h1 class="curproject-name"> u报名 </h1> 
 


# 组织

## 更改部门配置
<a id=更改部门配置> </a>
### 基本信息

**Path：** /org/dep/:id

**Method：** PATCH

**接口描述：**
<p>需要token<br>
名字不能重复且不能和现在的一样<br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"名字非法，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>
<p><span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1002</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"密码格式错误"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>
<p><span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"账号重复，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| id |   |  部门id<br/> |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 组织删除部门
<a id=组织删除部门> </a>
### 基本信息

**Path：** /org/dep/:id

**Method：** DELETE

**接口描述：**


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| id |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody">
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">删除的部门id</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 组织招新数据
<a id=组织招新数据> </a>
### 基本信息

**Path：** /org/get-data

**Method：** GET

**接口描述：**
<p>如果出现<br>
{<br>
&nbsp; &nbsp; "code":1001,<br>
&nbsp; &nbsp; "msg":"系统数据出错，请联系后台人员"<br>
}<br>
就说明后台数据被篡改了，需要联系数据库相关的人员处理</p>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> org</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> Id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sum</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pass</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> male</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> female</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> dep</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> Id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sum</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> pass</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> male</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> female</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 组织注册
<a id=组织注册> </a>
### 基本信息

**Path：** /org/reg

**Method：** POST

**接口描述：**
<p>无需token，给前后端组织注册使用<br>
账号和组织名字不能重复<br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"账号重复，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"组织名字非法，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 组织添加面试时间
<a id=组织添加面试时间> </a>
### 基本信息

**Path：** /org/add-time

**Method：** POST

**接口描述：**
<p>需要token</p>
<p><span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1002</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"开始时间不能大于结束时间"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span><br>
不能设置昨天面试结束<br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1002</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"结束时间不能小于当前时间"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> start</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">开始的时间戳（s）</span></td><td key=5><p key=5><span style="font-weight: '700'">mock: </span><span>100000</span></p></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> end</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">结束的时间戳（s）</span></td><td key=5><p key=5><span style="font-weight: '700'">mock: </span><span>90000000</span></p></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> msg</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 组织登录
<a id=组织登录> </a>
### 基本信息

**Path：** /org/login

**Method：** POST

**接口描述：**
<p>不需token</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessToken</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">token登录凭证</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessExpire</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">过期时间</span></td><td key=5></td></tr><tr key=0-4><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> refreshAfter</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">建议客户端刷新时间</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取组织下全部部门
<a id=获取组织下全部部门> </a>
### 基本信息

**Path：** /org/deps

**Method：** GET

**接口描述：**
<p>需要token</p>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> departments</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> account</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取组织下单个部门信息
<a id=获取组织下单个部门信息> </a>
### 基本信息

**Path：** /org/dep/:id

**Method：** GET

**接口描述：**
<p>需要token<br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1002</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"组织下无此部门"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| id |   |  部门id |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">部门id</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取组织招新时间
<a id=获取组织招新时间> </a>
### 基本信息

**Path：** /org/enroll-time/{:orgId}

**Method：** GET

**接口描述：**
<p>无需 token</p>


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| :orgId |   |  组织 id |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> start</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">开始时间戳</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> end</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">结束时间戳</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 部门注册
<a id=部门注册> </a>
### 基本信息

**Path：** /org/cre-dep

**Method：** POST

**接口描述：**
<p>需要组织登录token，格式为：</p>
<pre><code>'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
    -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI4NjcwNzQsImlhdCI6MTYxMjc4MDY3NCwidXNlcklkIjoxfQ.JKa83g9BlEW84IiCXFGwP2aSd0xF3tMnxrOzVebbt80'
</code></pre>
<p><span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"部门名字非法，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1006</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"密码不合法，请检查并修改"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">{</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"code"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(181, 206, 168)">1005</span><span class="colour" style="color:rgb(220, 220, 220)">,</span><br>
<span class="colour" style="color:rgb(212, 212, 212)">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color:rgb(156, 220, 254)">"msg"</span><span class="colour" style="color:rgb(220, 220, 220)">:</span><span class="colour" style="color:rgb(212, 212, 212)">&nbsp;</span><span class="colour" style="color:rgb(206, 145, 120)">"账号重复，请更换"</span><br>
<span class="colour" style="color:rgb(220, 220, 220)">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> account</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> password</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# 部门

## 删除面试官
<a id=删除面试官> </a>
### 基本信息

**Path：** /dep/delete-Inter/:uid

**Method：** GET

**接口描述：**


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| uid |   |   |

## 批量通过面试
<a id=批量通过面试> </a>
### 基本信息

**Path：** /dep/pass/:num

**Method：** POST

**接口描述：**
<p>需要部门登录token</p>
<p>如果num给错了：</p>
<p><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1002</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"面试只有一轮和二轮"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| num |  1/2 |  表示通过第一轮面试还是第二轮面试 |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">uid是一个int类型的数组</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>number</span></p></td></tr><tr key=array-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> msg</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=5><span style="font-weight: '700'">mock: </span><span>提交成功</span></p></td></tr>
               </tbody>
              </table>
            
## 查看录取名单
<a id=查看录取名单> </a>
### 基本信息

**Path：** /dep/getenroll

**Method：** POST

**接口描述：**
<p>参照前一接口</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/x-www-form-urlencoded | 是  |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody">
               </tbody>
              </table>
            
## 查看淘汰名单
<a id=查看淘汰名单> </a>
### 基本信息

**Path：** /dep/getout

**Method：** POST

**接口描述：**
<p>错误处理与上个接口相同。</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> page</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> num</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 查看第1面所有已面试用户的状态
<a id=查看第1面所有已面试用户的状态> </a>
### 基本信息

**Path：** /dep/getusers-first/interviewed

**Method：** POST

**接口描述：**
<p>假设只有一页的数量，你请求第二页的数据，就会报错：<br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"code"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168); --darkreader-inline-color:#e5e3be;" data-darkreader-inline-color="">1002</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"msg"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120); --darkreader-inline-color:#d8b79b;" data-darkreader-inline-color="">"超过数据总量"</span><br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">}</span></p>
<p><span class="colour" style="color: rgb(255, 249, 222); --darkreader-inline-color:#ffffe4;" data-darkreader-inline-color="">需要注意的就是如果返回的sum是0，则不会产生错误，不管你发送什么页数都是返回sum:0。</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color=""></span></p>
<p><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">如果报错的code为1001，说明数据库数据出错了，联系后台相应人员</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> page</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">页数，page从0开始，默认为0</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> num</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">每页显示数量，默认为10</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sum</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数据总量，方便前端分页</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> users</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> uid</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sex</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> phone</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> student_num</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> intro</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-7><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> second_status</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-8><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> final_status</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-9><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> eva</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-10><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> mark</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 添加面试官
<a id=添加面试官> </a>
### 基本信息

**Path：** /dep/add-inter/:uid

**Method：** POST

**接口描述：**
<p>uid为不存在的：<br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"code"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168); --darkreader-inline-color:#e5e3be;" data-darkreader-inline-color="">1004</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"msg"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120); --darkreader-inline-color:#d8b79b;" data-darkreader-inline-color="">"用户不存在"</span><br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">}</span></p>
<p>重复添加面试官：<br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"code"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168); --darkreader-inline-color:#e5e3be;" data-darkreader-inline-color="">1005</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#dcece5;" data-darkreader-inline-color="">"msg"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#f8e9d0;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120); --darkreader-inline-color:#d8b79b;" data-darkreader-inline-color="">"面试官已存在"</span><br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#fff0d6;" data-darkreader-inline-color="">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/x-www-form-urlencoded | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| uid |   |   |

## 用户被录取
<a id=用户被录取> </a>
### 基本信息

**Path：** /dep/enroll

**Method：** POST

**接口描述：**
<p>无论传入的uid对不对都只会返回录取成功，不对的uid数据库不会进行操作</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">数组，支持多个uid</span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>number</span></p></td></tr><tr key=array-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> </span></td><td key=1><span></span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody">
               </tbody>
              </table>
            
## 用户被淘汰
<a id=用户被淘汰> </a>
### 基本信息

**Path：** /dep/out

**Method：** POST

**接口描述：**
<p>参照用户被录取接口，参数和返回一样的处理</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/x-www-form-urlencoded | 是  |   |   |

## 获取一面所有人的状态
<a id=获取一面所有人的状态> </a>
### 基本信息

**Path：** /dep/getusers-first

**Method：** POST

**接口描述：**
<p>need token</p>
<p>逻辑上是必须填写简历才能报名，如果没填写简历就报名就是非法数据，如果由于恶意攻击或后台人员误操作数据库，会返回：<br>
意思是报名的uid里面有没填简历的，这儿是非法的<br>
<span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1001</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"有非法数据，请联系后台人员"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> page</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认为0，代表第page+1页</span></td><td key=5><p key=5><span style="font-weight: '700'">mock: </span><span>0</span></p></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> num</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">默认为10，代表一页有几个数据</span></td><td key=5><p key=5><span style="font-weight: '700'">mock: </span><span>10</span></p></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sum</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">人数总数</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> users</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> uid</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sex</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> phone</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> student_num</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> intro</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-7><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> status</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取部门报名信息
<a id=获取部门报名信息> </a>
### 基本信息

**Path：** /dep/reg-data

**Method：** POST

**接口描述：**
<p>需要token</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> reg_sum</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> pass_sum</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> pass_male</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> pass_female</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 部门登录
<a id=部门登录> </a>
### 基本信息

**Path：** /dep/login

**Method：** GET

**接口描述：**


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody">
               </tbody>
              </table>
            
# 用户

## 查看简历是否存在/查看简历
<a id=查看简历是否存在/查看简历> </a>
### 基本信息

**Path：** /user/check-resume

**Method：** GET

**接口描述：**
<p>需要token</p>
<p>注意，如果exist为false则是没提交过简历，这时候拿到的信息就为空了：<br>
<span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"exist"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">false</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"resume"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"name"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"sex"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"student_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"address_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"major"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"phone_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"email"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"intro"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">""</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(220, 220, 220);">}</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>
<p><span class="colour" style="color: rgb(220, 220, 220);">正常应该是：</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"exist"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">true</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"resume"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"name"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"lll"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"sex"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"男"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"student_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"20009288"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"address_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"竹一"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"major"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"计算机"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"phone_num"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"15525555555"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"email"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"<a href="mailto:1721261216@qq.com">1721261216@qq.com</a>"</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"intro"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"test"</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(220, 220, 220);">}</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> exist</span></td><td key=1><span>boolean</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> resume</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> sex</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> student_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> address_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> major</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> phone_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1-7><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> intro</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 用户填写简历
<a id=用户填写简历> </a>
### 基本信息

**Path：** /user/resume

**Method：** POST

**接口描述：**
<p>需要用户token<br>
手机和邮箱需要符合正则表达</p>
<p><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1002</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"手机号格式错误"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>
<p><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1002</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"邮箱格式错误"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sex</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> student_num</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> address_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-4><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phone_num</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-5><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> email</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-6><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> major</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-7><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> intro</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> msg</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">第一次是提交成功，第二次以后是修改简历成功</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 用户报名
<a id=用户报名> </a>
### 基本信息

**Path：** /user/register

**Method：** POST

**接口描述：**
<p>需要用户token</p>
<p><span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1005</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"你已经报名过青协test1组织的test3部门了"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> org_id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">组织id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> dep_id</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">部门</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> msg</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">报名成功</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 用户登录
<a id=用户登录> </a>
### 基本信息

**Path：** /user/login/:code

**Method：** GET

**接口描述：**
<p>无需token<br>
<span class="colour" style="color: rgb(220, 220, 220);">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"code"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168);">1002</span><span class="colour" style="color: rgb(220, 220, 220);">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212);">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254);">"msg"</span><span class="colour" style="color: rgb(220, 220, 220);">:</span><span class="colour" style="color: rgb(212, 212, 212);">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120);">"登录code不合法"</span><br>
<span class="colour" style="color: rgb(220, 220, 220);">}</span></p>


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| code |  023xCIkl2NwDd84Wbhml2Zi6Wf0xCIkA |  前端wx.login拿到的code |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessToken</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户token</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessExpire</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">过期时间</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> refreshAfter</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">建议客户端刷新时间</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">主页的uid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 用户获取自己的报名情况
<a id=用户获取自己的报名情况> </a>
### 基本信息

**Path：** /user/register

**Method：** GET

**接口描述：**
<p>需要的token</p>
<p>first\second\final 里面的数字对应的状态：</p>
<pre><code>const (
   NEW         = 0 + iota // 0-新投递 
   ARRANGED               // 1-已经安排面试
   INTERVIEWED            // 2-已面试
   PASS                   // 3-通过
   OUT                    // 4-淘汰
)
</code></pre>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> rsp</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> dep_id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> dep_name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> org_id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> org_name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-4><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> first</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">第一面状态</span></td><td key=5></td></tr><tr key=0-0-5><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> second</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">二面状态</span></td><td key=5></td></tr><tr key=0-0-6><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> final</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">最终状态</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取正在招新的组织
<a id=获取正在招新的组织> </a>
### 基本信息

**Path：** /user/org-enrolling

**Method：** GET

**接口描述：**
<p>无需token</p>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> organizations</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> org_id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> org_name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-2><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> start_time</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr><tr key=0-0-3><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> end_time</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获得某个正在招新的组织下的所有部门
<a id=获得某个正在招新的组织下的所有部门> </a>
### 基本信息

**Path：** /user/org-deps/:id

**Method：** GET

**接口描述：**
<p>需要token</p>


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| id |   |  前面获取正在招新的组织那里给的组织id |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> deps</span></td><td key=1><span>object []</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap"></span></td><td key=5><p key=3><span style="font-weight: '700'">item 类型: </span><span>object</span></p></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> dep_id</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">部门id，报名的时候需要</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> dep_name</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">部门名字</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
# 面试官

## 查看评价
<a id=查看评价> </a>
### 基本信息

**Path：** /inter/vieweva/:num

**Method：** POST

**接口描述：**
<p>需要token</p>
<p>这个是num小于1 或者大于2<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"轮次错误"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; ”code":1004,<br>
&nbsp; &nbsp; "msg":"用户没有注册"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户还没有参加第一轮面试"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户还没有参加第二轮面试"<br>
}</p>
<p><span class="colour" style="color: rgb(181, 174, 162); --darkreader-inline-color:#b6afa3;" data-darkreader-inline-color="">这个不是没有携带token错误 这个是面试官所能面试的部门和用户报名的部门不一样</span><br>
<span class="colour" style="color: rgb(181, 174, 162); --darkreader-inline-color:#b6afa3;" data-darkreader-inline-color="">{</span><br>
<span class="colour" style="color: rgb(181, 174, 162); --darkreader-inline-color:#b6afa3;" data-darkreader-inline-color="">&nbsp; &nbsp; "code":1007,</span><br>
<span class="colour" style="color: rgb(181, 174, 162); --darkreader-inline-color:#b6afa3;" data-darkreader-inline-color="">&nbsp; &nbsp; "msg":"面试官没有权限"</span><br>
<span class="colour" style="color: rgb(181, 174, 162); --darkreader-inline-color:#b6afa3;" data-darkreader-inline-color="">}</span></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| num |  1 |  这轮面试的轮次 |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> depId</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的部门id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所要查看用户评价的用户uid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> eva</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户评价</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> mark</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">用户分数 1是A 2是B 3是C</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 给用户打分评价
<a id=给用户打分评价> </a>
### 基本信息

**Path：** /inter/evaluate/:num

**Method：** POST

**接口描述：**
<p>需要token<br>
这个是num小于1 或者大于2<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"轮次错误"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; ”code":1004,<br>
&nbsp; &nbsp; "msg":"用户没有注册"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户已经被淘汰了"<br>
}</p>
<p>用户面试过了 或者已经被录取了<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户无需参加此轮面试"<br>
}</p>
<p>这个是如果第二轮面试的时候发现第一轮面试都没过的人来参加了<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户没有参加第一轮面试"<br>
}</p>
<p>这个不是没有携带token错误 这个是面试官所能面试的部门和用户报名的部门不一样<br>
{<br>
&nbsp; &nbsp; "code":1007,<br>
&nbsp; &nbsp; "msg":"面试官没有权限"<br>
}</p>
<p><br data-tomark-pass=""><br>
<br data-tomark-pass=""></p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| num |  1 |  面试轮次 |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> depId</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的部门id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所要评价的用户的uid</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> eva</span></td><td key=1><span>string</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官对用户的评价</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> mark</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官对用户的打分  A是1 B是2 C是3</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> msg</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">返回 评价成功 这个字符串</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取用户的信息
<a id=获取用户的信息> </a>
### 基本信息

**Path：** /inter/getresume/:num

**Method：** POST

**接口描述：**
<p>需要token</p>
<p>这个是num小于1 或者大于2<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"轮次错误"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; ”code":1004,<br>
&nbsp; &nbsp; "msg":"用户不存在"<br>
}</p>
<p>{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户已经被淘汰了"<br>
}</p>
<p>这个是如果第二轮面试的时候发现第一轮面试都没过的人来参加了<br>
{<br>
&nbsp; &nbsp; "code":1002,<br>
&nbsp; &nbsp; "msg":"用户没有参加第一轮面试"<br>
}</p>
<p>这个不是没有携带token错误 这个是面试官所能面试的部门和用户报名的部门不一样<br>
{<br>
&nbsp; &nbsp; "code":1007,<br>
&nbsp; &nbsp; "msg":"面试官没有权限"<br>
}</p>


### 请求参数
**Headers**

| 参数名称  | 参数值  |  是否必须 | 示例  | 备注  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| Content-Type  |  application/json | 是  |   |   |
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| num |  1 |  当前所面试轮次 |
**Body**

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> depId</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的部门id</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number</span></td><td key=2>必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">所要获取用户信息的用户uid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> name</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">名字</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> sex</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">性别</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> student_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">学号</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> address_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">地址</span></td><td key=5></td></tr><tr key=0-4><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> major</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">专业</span></td><td key=5></td></tr><tr key=0-5><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> phone_num</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">电话号码</span></td><td key=5></td></tr><tr key=0-6><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> email</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">邮箱</span></td><td key=5></td></tr><tr key=0-7><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> intro</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">自我介绍</span></td><td key=5></td></tr><tr key=0-8><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> tag</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">标志位 1标志着这个用户uid重复输入过已经被这轮面试评价过了或者已经被选为部员只能查看评价了 0代表着这个用户还没有被评价过可以进行评价</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 获取面试官的组织和部门
<a id=获取面试官的组织和部门> </a>
### 基本信息

**Path：** /inter/gerorgdep

**Method：** GET

**接口描述：**
<p>需要token<br>
{<br>
&nbsp; &nbsp; "code":1007,<br>
&nbsp; &nbsp; "msg":"没有面试官权限"<br>
}<br>
id和名字都给是因为后面接口有需要面试官在这里所选择的id 后面接口需要传入</p>


### 请求参数

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> depId</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的部门id和对应的名字</span></td><td key=5></td></tr><tr key=0-0-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> 5</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">id:名字  前面的数字是id后面跟着名字 后同</span></td><td key=5></td></tr><tr key=0-0-1><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> 6</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">id:名字</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> orgId</span></td><td key=1><span>object</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">组织的id和对应的名字</span></td><td key=5></td></tr><tr key=0-1-0><td key=0><span style="padding-left: 20px"><span style="color: #8c8a8a">├─</span> 43</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">id:名字</span></td><td key=5></td></tr>
               </tbody>
              </table>
            
## 面试官登录
<a id=面试官登录> </a>
### 基本信息

**Path：** /inter/login/:code

**Method：** GET

**接口描述：**
<p>无需token<br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#d5d0c8;" data-darkreader-inline-color="">{</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#d0cac2;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#8dd6fa;" data-darkreader-inline-color="">"code"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#d5d0c8;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#d0cac2;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(181, 206, 168); --darkreader-inline-color:#b3cba2;" data-darkreader-inline-color="">1002</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#d5d0c8;" data-darkreader-inline-color="">,</span><br>
<span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#d0cac2;" data-darkreader-inline-color="">&nbsp;&nbsp;&nbsp;&nbsp;</span><span class="colour" style="color: rgb(156, 220, 254); --darkreader-inline-color:#8dd6fa;" data-darkreader-inline-color="">"msg"</span><span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#d5d0c8;" data-darkreader-inline-color="">:</span><span class="colour" style="color: rgb(212, 212, 212); --darkreader-inline-color:#d0cac2;" data-darkreader-inline-color="">&nbsp;</span><span class="colour" style="color: rgb(206, 145, 120); --darkreader-inline-color:#d0967c;" data-darkreader-inline-color="">"登录code不合法"</span><br>
<span class="colour" style="color: rgb(220, 220, 220); --darkreader-inline-color:#d5d0c8;" data-darkreader-inline-color="">}</span></p>


### 请求参数
**路径参数**

| 参数名称 | 示例  | 备注  |
| ------------ | ------------ | ------------ |
| code |  023xCIkl2NwDd84Wbhml2Zi6Wf0xCIkA |  面试官小程序给的code |

### 返回数据

<table>
  <thead class="ant-table-thead">
    <tr>
      <th key=name>名称</th><th key=type>类型</th><th key=required>是否必须</th><th key=default>默认值</th><th key=desc>备注</th><th key=sub>其他信息</th>
    </tr>
  </thead><tbody className="ant-table-tbody"><tr key=0-0><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessToken</span></td><td key=1><span>string</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的token</span></td><td key=5></td></tr><tr key=0-1><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> accessExpire</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">过期时间</span></td><td key=5></td></tr><tr key=0-2><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> refreshAfter</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">建议浏览器的刷新时间</span></td><td key=5></td></tr><tr key=0-3><td key=0><span style="padding-left: 0px"><span style="color: #8c8a8a"></span> uid</span></td><td key=1><span>number</span></td><td key=2>非必须</td><td key=3></td><td key=4><span style="white-space: pre-wrap">面试官的uid</span></td><td key=5></td></tr>
               </tbody>
              </table>
            