<template>
    <li>
        <div class="ti">{{item.title}}<i>{{item.time}}</i></div>
        <div class="con">{{item.content}}</div>
    </li>
</template>
<script>
export default {
    name: 'LeeNotificationItem',
    data() {
        return {
           
        }
    },   
    mounted() {
        
    },
    computed: {
        
    },
    props: {
        item: {
            type: Object,
            default () {
                return {}
            }
        }
    }
};
</script>
<style scoped>
ul,
li {
    list-style: none;
    margin: 0px;
    padding: 0px;
}

.lee-notification ul {
    position: relative;
}

.lee-notification ul li {
    font-size: 14px;
    background: rgba(255, 255, 255, .9);
    border-radius: 6px;
    margin-bottom: 10px;
    transition: all .3s;
}

.lee-notification ul li .ti {
    background: rgba(255, 255, 255, .9);
    padding: 5px 10px;
    border-radius: 6px 6px 0 0;
    color: #999
}

.lee-notification ul li .ti i {
    float: right;
    color: #999
}

.lee-notification ul li .con {
    line-height: 22px;
    padding: 10px;
}


.lee-enter {
    opacity: 0;
    transform: scale(.2);
}

.lee-enter-to {
    opacity: 1;
    transform: scale(1);
}

.lee-leave {
    opacity: 1;
}

.lee-leave-to {
    opacity: 0;
}

</style>