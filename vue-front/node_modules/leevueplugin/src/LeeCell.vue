<template>
    <div class="lee-cell">
        <div class="lee-cell-icon" v-if="icon.length>0"><img :src="icon"></div>
        <div class="lee-cell-title">{{title}}</div>
        <div class="lee-cell-footer" :class="{del:del,link:!del}"  @click="goto(url)">{{footer}}</div>
    </div>
</template>
<script>
export default {
    name: 'LeeCell',
    data() {
        return {

        }
    },
    methods: {
        goto(url){
            if(this.del){
                this.$emit('del')
                return
            }
            if(url.length>0){
                location.href=url
            }

        }

    },
    props: {
        title: {
            type: String,
            default: 'Cells的标题' //默认default
        },
        footer: {
            type: String,
            default: 'Cells底部的文字' //默认default
        },
        url: {
            type: String,
            default: '' //默认default
        },
        icon: {
            type: String,
            default: '' //默认default
        },
        del: {
            type: Boolean,
            default: false
        }
    }
};
</script>
<style scoped>


.lee-cell {
    cursor: pointer;
    display: flex;
    padding: 16px;
    width: 100%;
    border-bottom: solid 1px #dedede;box-sizing: border-box;
    align-items: center
}
.lee-cell:last-child{border-bottom: none}
.lee-cell-title {
    color: rgba(0, 0, 0, .8);
    flex-grow: 1
}
.lee-cell-icon{width: 26px;height: 26px;padding-right: 15px}

.lee-cell-footer {
    color: rgba(0, 0, 0, .4);
    position: relative;
}

.link:after{content: '>';color: #ddd;padding-left: 14px}
.del:after{content: 'x';color: #ddd;padding-left: 14px}
</style>
