<template>
    <transition name="fade">
        <div :class="['lee-alert','lee-alert-'+type]" v-if="isShow">
            <i class="lee-alert-icon" v-if="showIcon"></i>
            <div class="lee-alert-content">{{title}}</div>
            <i class="lee-alert-close" v-if="closable" @click="hide"></i>
        </div>
    </transition>
</template>
<script>
export default {
    name: 'LeeAlert',
    data() {
        return {
            isShow: true

        }
    },
    watch: {
        isShow: {
            immediate: true,
            handler(value) {
                value = this.closable

            }
        }
    },
    methods: {
        hide(){
            this.isShow=false;
            this.$emit('close')
        }
    },
    props: {
        title: {
            type: String,
            default: '' //默认default
        },
        type: {
            type: String,
            default: 'warning' //默认default
        },
        closable: {
            type: Boolean,
            default: true //默认default
        },
        showIcon: {
            type: Boolean,
            default: false //默认default
        }
    }
};
</script>
<style scoped>
.lee-alert {
    padding: 8px 16px;
    box-sizing: border-box;
    border-radius: 4px;
    position: relative;
    background-color: #fff;
    overflow: hidden;
    opacity: 1;
    display: flex;
    align-items: center;
    transition: opacity .2s;
    margin:15px 0;
}
.lee-alert-icon{width: 16px;height: 16px;margin-right: 8px;border-radius: 50%;color: #fff;text-align: center;font-style: normal;position: relative;}
.lee-alert-icon:before{font-size: 16px;height: 16px;line-height: 16px;display: block}
.lee-alert-success .lee-alert-icon{background-color: #67c23a}
.lee-alert-success .lee-alert-icon:before{content: "√";}
.lee-alert-info .lee-alert-icon{background-color: #909399}
.lee-alert-info .lee-alert-icon:before{content: "i";}
.lee-alert-warning .lee-alert-icon{background-color: #e6a23c}
.lee-alert-warning .lee-alert-icon:before{content: "!";}
.lee-alert-error .lee-alert-icon{background-color: #f56c6c}
.lee-alert-error .lee-alert-icon:before{content: "x";line-height: 12PX;}
.lee-alert-success {
    background-color: #f0f9eb;
    color: #67c23a
}

.lee-alert-info {
    background-color: #f4f4f5;
    color: #909399
}

.lee-alert-warning {
    background-color: #fdf6ec;
    color: #e6a23c
}

.lee-alert-error {
    background-color: #fef0f0;
    color: #f56c6c
}

.lee-alert .lee-alert-content {
    font-size: 13px;
    flex-grow: 1
}

.lee-alert .lee-alert-close {
    width: 12px;
    height: 12px;
    position: relative;
    transform: rotate(45deg);
    cursor: pointer;
}

.lee-alert .lee-alert-close:before {
    content: "";
    width: 100%;
    height: 1px;
    background: #ccc;
    position: absolute;
    top: 50%;
    margin-top: -0.5px
}

.lee-alert .lee-alert-close:after {
    content: "";
    width: 1px;
    height: 100%;
    background: #ccc;
    position: absolute;
    left: 50%;
    margin-left: -0.5px
}
.fade-leave-active {
    transition: all .4s;
}

.fade-leave-to {
    opacity: 0;
}
</style>