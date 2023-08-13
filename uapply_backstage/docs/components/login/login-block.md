# Login Block

## Where to Use ?

You should use this component in the login page.


## For What ?

To pull the input logic and login block out.


## Path

> src/components/login/login_block.vue


## Props

| Name | Type | Instruction |
|------|------|-------------|
| /    |   /  |   / |


## Emits

| Name | Props | Instruction | Trigger |
|------|------|--------------|---------|
| login |  Object(role,account,password)  | Login block return user information to the login page | When click login button|


## Defined Variable

| Name | Type | Instruction | Required |
|------|------|-------------|----------|
| block_title | string | the title of the login block | true |
| isPassword | boolean | switch the type of password input label | true |
| isVisible | boolean | switch the type of eye image | true |
| role_options | object | the n-select options | true |
| account_infor | object | store the information of user's input | true |


## Events

| Name | Instruction | Callback Parameter |
|------|-------------|--------------------|
| changeVisible | switch the visibility of the password when click eye image | / |
| login | return the user input to login page when click login button | / |
| keyup | when 'enter' is pressed up,use function login to login in | / |


## Need Adaptation

| Where | What | Supplement | 
|-------|------|------------|
| Login button | Button text "登录" | Now it's seems to be indifferent |
| Login title | Login block text "U报名后台管理系统-登录" | Now it's seems to be indifferent |