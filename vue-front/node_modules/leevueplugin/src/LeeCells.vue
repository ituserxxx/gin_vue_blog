<template>
    <div class="lee-cells">
        <div class="lee-cells-title" v-if="title.length>0">{{title}}</div>
        <div class="lee-cells-inner">
            <slot></slot>
        </div>
        <div class="lee-cells-footer" v-if="footer.length>0">{{footer}}</div>
    </div>
</template>
<script>
export default {
    name: 'LeeCells',
    data() {
        return {

        }
    },
    methods: {

    },
    props: {
        title: {
            type: String,
            default: '' //默认default
        },
        footer: {
            type: String,
            default: '' //默认default
        }
    }
};
</script>
<style scoped>
.lee-cells-title,
.lee-cells-footer {
    font-size: 14px;
    color: rgba(0, 0, 0, .5);
}

.lee-cells-title {
    padding-bottom: 6px;
}

.lee-cells-inner {
    background: #fff;
    font-size: 16px;
    border: solid 1px #dedede;
    display: flex;
    align-items: flex-end;
    flex-direction: column;
}

.lee-cells-footer {
    padding-top: 6px;
}
</style>