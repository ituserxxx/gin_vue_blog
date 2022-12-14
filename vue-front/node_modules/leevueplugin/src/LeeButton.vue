<template>
    <button type="button" :class="['lee-button','lee-button-'+type]" v-on="buttonListeners">
        <span><slot></slot></span>
    </button>
</template>
<script>
export default {
    name: 'LeeButton',
    data() {
        return {

        };
    },
    props: {
        type: {
            type: String,
            default: 'default' //默认default
        }
    },
    computed: {
        buttonListeners() {
            var vm = this;
            return Object.assign({}, this.$listeners, {
                click: function(event) {
                    vm.$emit('click', event);
                }
            })
        }
    }
};
</script>
<style scoped>
.lee-button {
    display: inline-block;
    line-height: 1;
    white-space: nowrap;
    cursor: pointer;
    background: #fff;
    border: 1px solid #dcdfe6;
    border-color: #dcdfe6;
    color: #606266;
    -webkit-appearance: none;
    text-align: center;
    box-sizing: border-box;
    outline: none;
    margin: 0;
    transition: .1s;
    font-weight: 500;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    padding: 12px 20px;
    font-size: 14px;
    border-radius: 4px;
}

.lee-button-primary {
    color: #fff;
    background-color: #409eff;
    border-color: #409eff;
}

.lee-button-success {
    color: #fff;
    background-color: #67c23a;
    border-color: #67c23a;
}

.lee-button-info {
    color: #fff;
    background-color: #909399;
    border-color: #909399;
}

.lee-button-warning {
    color: #fff;
    background-color: #e6a23c;
    border-color: #e6a23c;
}

.lee-button-danger {
    color: #fff;
    background-color: #f56c6c;
    border-color: #f56c6c;
}
</style>