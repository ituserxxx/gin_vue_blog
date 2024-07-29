<template>
    <div class="lee-steps" :class="[direction,fillstyle,centerstyle]">
       <slot></slot>
    </div>
</template>
<script>
export default {
    name: 'LeeSteps',
    inheritAttrs: false,
    data() {
        return {
          
        }
    }, 
    mounted(){

    },
    computed:{
        fillstyle(){
            if(this.fill){
                return 'fill'
            }
        },
        centerstyle(){
            if(this.center){
                return 'center'
            }else{
                return 'normal'
            }
        },
    }
    ,
    provide() {
        return {
            step: this
        };
    },
    props: {
        active: {
            type: Number,
            default: 0 //默认default
        },
        color: {
            type: String,
            default: '#46bd87' //默认default
        },
         fill: {
            type: Boolean,
            default: false //默认default
        },
         center: {
            type: Boolean,
            default: false //默认default
        },
         finish: {
            type: String,
            default: 'finish' //默认default
        },
        direction: {
            type: String,
            default: 'horizontal' //默认default  vertical 
        }

    },
    
};
</script>
<style scoped>
.lee-steps{display: flex;}
.lee-steps.vertical{flex-direction: column;height: 100%;}
.lee-steps.vertical .lee-step{display: flex;}

</style>