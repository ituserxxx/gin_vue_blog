<template>
    <transition name="fade">
        <div :class="['lee-alert','lee-alert-'+thetype]" v-if="isShow">
            <i class="lee-alert-icon" v-if="showIcon"></i>
            <div class="lee-alert-content">{{title}}</div>
            <i class="lee-alert-close" v-if="closeable" @click="hide"></i>
        </div>
    </transition>
</template>
<script>
export default {
    name: 'LeeAlert',
    data() {
        return {
            isShow: true,
            thetype:'warning'

        }
    },
    watch: {
        type: {
            immediate: true,
            handler(value) {
                let types=['primary','success','danger','info','warning','leecolor']
                let fi=types.filter(function(el) {
                    return el.indexOf(value)>-1;
                })
                if(fi.length<1){
                    this.thetype="info"
                }else{
                    this.thetype=value
                }
            }
        },
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
            default: '提示文字' //默认default
        },
        type: {
            type: String,
            default: 'info' //默认default
        },
        closeable: {
            type: Boolean,
            default: false //默认default
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
.lee-alert-leecolor .lee-alert-icon{background-color: #46bd87}
.lee-alert-leecolor .lee-alert-icon:before{content: "√";}
.lee-alert-success .lee-alert-icon{background-color: #67c23a}
.lee-alert-success .lee-alert-icon:before{content: "√";}
.lee-alert-info .lee-alert-icon{background-color: #909399}
.lee-alert-info .lee-alert-icon:before{content: "i";}
.lee-alert-warning .lee-alert-icon{background-color: #e6a23c}
.lee-alert-warning .lee-alert-icon:before{content: "!";}
.lee-alert-danger .lee-alert-icon{background-color: #f56c6c}
.lee-alert-danger .lee-alert-icon:before{content: "x";line-height: 12PX;}
.lee-alert-primary .lee-alert-icon{background-color: #409eff}
.lee-alert-primary .lee-alert-icon:before{content: "i";}
.lee-alert-leecolor {
    background-color: rgba(70,189,135,.1);
    color: #46bd87
}
.lee-alert-success {
    background-color:rgba(103,194,58,.1);
    color: #67c23a
}

.lee-alert-info {
    background-color: #f4f4f5;
    color: #909399
}

.lee-alert-warning {
    background-color: rgba(230,162,60,.1);
    color: #e6a23c
}

.lee-alert-danger {
    background-color: rgba(245,108,108,.1);
    color: #f56c6c
}
.lee-alert-primary {
    background-color: rgba(64,158,255,.1);
    color: #409eff
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
