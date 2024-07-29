<template>
  <a :href="href">{{title}}<span :style="[styleObject]">{{info}}</span></a>
</template>
<script>
export default {
    name:'LeeBadge',
    data() {
        return {
         styleObject:{
            background:this.color
         }
        }
    },
    props: {
        title: {
            type: String,
            default: '徽章名称' //默认default
        },
        info: {
            type: String,
            default: '0' //默认default
        },
        color: {
            type: String,
            default: '#f00' //默认default
        },
        href: {
            type: String,
            default: '' //默认default
        }
    }
};
</script>
<style scoped>
a{margin:0 8px;}
a span{font-size: 12px;background: #f44;border-radius: 4px;color: #fff;position: relative;text-align: center;min-width: 14px;top: -4px;display: inline-block;line-height: 14px;}
</style>