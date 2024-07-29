<template>
    <div class="lee-radio-group" :class="[size,type]" ><slot></slot></div>
</template>
<script>
export default {
    name: 'LeeRadioGroup',
    inheritAttrs:false,
    data() {
        return {
        };
    },
    provide() {
        return {
            radio: this
        };
    },
    props: {
        value:{
            type:Number,
            default:1
        },
        size:{
            type:String,
            default:'s'
        },
        type:{
            type:String,
            default:''
        }
    },
    computed: {
        
    }
};
</script>
<style scoped>
.lee-radio-group{display:flex;align-items: center}
.lee-radio-group.s .lee-radio{height: 16px;line-height: 16px;font-size: 16px;padding-left: 24px;}
.lee-radio-group.s .lee-radio:before{width: 16px;height: 16px}
.lee-radio-group.s .lee-radio:after{width: 10px;height: 10px}
.lee-radio-group.m .lee-radio{height: 20px;line-height: 20px;font-size: 20px;padding-left: 28px;}
.lee-radio-group.m .lee-radio:before{width: 20px;height: 20px}
.lee-radio-group.m .lee-radio:after{width: 14px;height: 14px}
.lee-radio-group.l .lee-radio{height: 24px;line-height: 24px;font-size: 24px;padding-left: 32px;}
.lee-radio-group.l .lee-radio:before{width: 24px;height: 24px}
.lee-radio-group.l .lee-radio:after{width: 18px;height: 18px}
.lee-radio-group.primary .lee-radio.check:before{border: solid 1px #409eff;}
.lee-radio-group.primary .lee-radio.check:after{background: #409eff}
.lee-radio-group.primary .lee-radio:hover:before{border: solid 1px #409eff;}
.lee-radio-group.success .lee-radio.check:before{border: solid 1px #67c23a;}
.lee-radio-group.success .lee-radio.check:after{background: #67c23a}
.lee-radio-group.success .lee-radio:hover:before{border: solid 1px #67c23a;}
.lee-radio-group.info .lee-radio.check:before{border: solid 1px #909399;}
.lee-radio-group.info .lee-radio.check:after{background: #909399}
.lee-radio-group.info .lee-radio:hover:before{border: solid 1px #909399;}
.lee-radio-group.warning .lee-radio.check:before{border: solid 1px #e6a23c;}
.lee-radio-group.warning .lee-radio.check:after{background: #e6a23c}
.lee-radio-group.warning .lee-radio:hover:before{border: solid 1px #e6a23c;}
.lee-radio-group.danger .lee-radio.check:before{border: solid 1px #f56c6c;}
.lee-radio-group.danger .lee-radio.check:after{background: #f56c6c}
.lee-radio-group.danger .lee-radio:hover:before{border: solid 1px #f56c6c;}
</style>