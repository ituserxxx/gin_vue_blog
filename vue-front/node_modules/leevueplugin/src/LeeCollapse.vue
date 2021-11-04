<template>
    <div class="lee-collapse">
        <slot></slot>
    </div>
</template>
<script>
export default {
    name: 'LeeCollapse',
    inheritAttrs: false,
    data() {
        return {
           expname:''
        }
    },
    provide() {
        return {
            collapse: this
        };
    },
    watch: {
        value: {
            immediate: true,
            handler(value) {
                this.expname=value
                this.$emit('change',value)
            }
        }
    },
    props: {
        value:[String,Array],
        accordion: {
            type: Boolean,
            default:false
        }
    }
};
</script>
<style scoped>

</style>