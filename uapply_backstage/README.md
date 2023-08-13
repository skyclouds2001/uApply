# u_rigester


## What is u_rigester ? 

This is a background management system of wechat applet called U Rigester .
And this is the version 2.0 for the system.
The reason we have to rebuild this project is,the version 1.0 is too humble in some aspects.

The main problems in the previous version are :
1. The adaptation for screems with different resolutions is't very good.
2. There are some logical error or some logical loophole.
3. The previous workflow is too humble to complete a complete project in some aspects like log and so on.
4. The version 1.0 does have some deficiencies and some framework issues to be solved.


## Tool and Edition

If you have already installed a node , you can use a tool called NVM to help you to manage different node versions.

| tool name | corresponding edition |
|------------|--------------------|
|    VUE     | 3.0.0 |
| vue-router | 4.0.0 |
|   VUEX     | 4.0.0 |
|   Naive    | 2.31.0 |
|   node.js  | 14.17.6 |
|   sass     | 1.26.5   |
| vue-cookies | 1.8.1   |


## Some Project Operation

| operation name                            | operation instructions|
|------------------------------------------|------------|
| project setup                            | npm install|
| Compiles and hot-reloads for development | npm run serve|
| Compiles and minifies for production     | npm run build |
| Lints and fixes files                    | npm run lint|


## Menu Structure

Here is the menu structure of this project.
In the project development stage , thie menu is showed for the developer.
And there is some notes behind each line.You must read and get it.

For simplifing the expression,I used some symbols.And the correspondence is in the below table:

| expression | symbols |
|------------|---------|
| You should know its usage but it's not important | A |
| You should know its usage and master it | B |
| It's not necessary to care about it | C |
| You should focus on it ! | * |

It is necessary to mention that some files but not folder are not marked when some detail documents are on the docs folder.

If you want to learn more about each file's function,go [VUE Project Menu Structure](https://blog.csdn.net/weixin_44161385/article/details/123707269?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522165780401016780357238179%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=165780401016780357238179&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_click~default-2-123707269-null-null.142^v32^pc_search_result_control_group,185^v2^control&utm_term=vue%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84&spm=1018.2226.3001.4187)

```
|-- u_rigester 
    |-- .gitignore                  C
    |-- babel.config.js             C
    |-- package-lock.json           A
    |-- package.json                B : Here to configure the project
    |-- README.md                   A
    |-- vue.config.js               C
    |-- public
    |   |-- favicon.ico             C
    |   |-- index.html              C
    |-- src                         B : Main folder for most code and documents
        |-- App.vue                 B : Main component
        |-- main.js                 B : Entry file 
        |-- assets                  B : Static resources
        |   |-- logo.png            C
        |   |-- imgs                B : Image resources
        |-- components             *B*: All the components should be placed here 
        |   |-- HelloWorld.vue
        |-- docs                   *B*: All the documents are placed here
        |   |-- components         *B*: Documents of components
        |   |-- pages              *B*: Documents of pages
        |   |-- problems           *B*: If you have some problems,write here
        |-- router                  B : Used for VUE-Router
        |   |-- index.js
        |-- store                   B : Used for VUEX
        |   |-- index.js
        |-- views                  *B*: You will put your pages here
            |-- About.vue
            |-- Home.vue
```

## Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).


## Development Calendar

| date | have done |
|------|-----------|
| 2022.7.14 | project created and configed |