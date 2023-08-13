<template>
    <div class="department">
        <div class="top">
            <span>{{ "ID : " + info.id }}</span>
            <img :src="delete_img" @click="dep_delete" title="删除部门">
        </div>
        <div class="mid">
            <img :src="name_img" @click="dep_copy" title="复制部门信息">
            {{ info.name }}
        </div>
        <div class="bot">
            <div class="user">
                <span>{{ "账号：" + info.account }}</span>
                <span>{{ "密码：" + info.password }}</span>
            </div>
            <img :src="edit_img" @click="dep_edit" title="重新配置部门">
        </div>
    </div>
</template>

<script>
import { message } from '../../../utils/message/message';
import { getCurrentInstance } from 'vue';

export default {
    props: {
        // 当前部门的信息
        info: Object
    },
    setup(props, context) {
        const delete_img = require("@/assets/imgs/错误.png");
        const name_img = require("@/assets/imgs/复制.png");
        const edit_img = require("@/assets/imgs/编辑.png");

        // 实例对象
        const { proxy } = getCurrentInstance();

        // 删除部门
        const dep_delete = function () {
            proxy.$sessionData("set","ID",props.info.id);
            context.emit("toast", "DeleteDep");
        }
        // 编辑部门信息
        const dep_edit = function () {
            // console.log(props.info);
            proxy.$sessionData("set","department",props.info);
            context.emit("toast", "EditDep");
        }
        // 复制信息到剪切板
        const dep_copy = function () {
            copy("ID:" + props.info.id + "\n" + "部门:" + props.info.name + "\n" + "账号:" + props.info.account + "\n" + "密码:" + props.info.password);
        }

        // 复制文本到剪切板
        // $contain : 需要复制到剪切板的信息
        const copy = function (contain) {
            // 1、创建一个文本域节点
            var textArea = document.createElement("textarea");
            // 2、把要复制的内容插入到上面创建的文本域节点里
            textArea.textContent = contain;
            // 3、设置创建的文本域节点style
            textArea.style.position = "fixed";
            // 4、把创建的文本节点插入到body里
            document.body.appendChild(textArea);
            // 5、选择创建的文本域节点
            textArea.select();
            // 6、try { } catch () { } 捕获异常
            try {
                document.execCommand("copy"); // 把要复制的内容拷贝到剪贴板
                message.success("已复制部门信息到剪切板！");
            } catch (ex) {
                message.error("复制部门信息失败！");
                return false;
            } finally {
                document.body.removeChild(textArea); // 移除插入的文本域节点
            }
        }

        // 这里不能删，但是我也不想知道哪儿出错了
        const reset = function () {
        }

        return {
            delete_img,
            name_img,
            edit_img,
            dep_delete,
            dep_edit,
            dep_copy,
            reset
        }
    }
}
</script>

<style lang="scss" scoped>
img {
    cursor: pointer;
}

.department {
    border-radius: 0.5rem;
    background: white;
    margin: 3% 5% 0 0rem;
    box-shadow: 1px 1px 5px 0px #c5a0ff;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;

    height: 12rem;
    width: 22rem;

    @media (max-width:1440px) {
        width: 17.5rem;
        height: 9rem;
    }

    @media (max-width:1280px) {
        width: 15.5rem;
        height: 9rem;
    }

    .top {
        height: 22%;
        width: 85%;
        display: flex;
        justify-content: space-between;
        align-items: flex-end;

        span {
            color: #999999;
        }

        img {
            height: 90%;
            position: relative;
            top: 20%;
        }
    }

    .mid {
        height: 25%;
        width: 85%;
        font-weight: 550;

        font-size: 2rem;

        @media (max-width:1440px) {
            font-size: 1.6rem;
        }

        img {
            height: 80%;
            position: relative;
            top: 10%;
        }
    }

    .bot {
        height: 40%;
        width: 85%;
        margin-bottom: 5%;
        display: flex;
        justify-content: space-between;
        align-items: center;

        .user {
            height: 100%;
            display: flex;
            flex-direction: column;
            justify-content: space-around;

            font-size: 1.1rem;

            @media (max-width:1440px) {
                font-size: 0.9rem;
            }

            @media (max-width:1280px) {
                font-size: 0.8rem;
            }
        }

        img {
            height: 60%;
            position: relative;
            left: 2%;
            top: 15%;
        }
    }

}
</style>