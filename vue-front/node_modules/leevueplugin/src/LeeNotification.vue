<template>
    <div class="lee-notification">
        <transition-group name="lee" tag="ul">
            <slot></slot>
        </transition-group>
    </div>
</template>
<script>
export default {
    name: 'LeeNotification',
    data() {
        return {
        }
    }
};
</script>
<style scoped>
ul,li{
    list-style: none;
    margin: 0px;
    padding: 0px;
}


</style>