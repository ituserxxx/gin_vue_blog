<template>
    <div class="lee-tooltip" @mouseenter="enter"  @mouseleave="leave">
        <slot></slot>
        <transition :name="positionStyle">
            <div v-show="isshow" class="lee-tooltip-item" :class="positionStyle">
                <div class="lee-tooltip-itemn" :style="'--color:'+color" v-if="positionStyle==='left'||positionStyle==='top'">
                    <div class="content">{{content}}</div>
                    <div class="jt"><i></i></div>
                </div>
                <div class="lee-tooltip-itemn" :style="'--color:'+color" v-if="positionStyle==='right'||positionStyle==='bottom'">
                     <div class="jt"><i></i></div>
                    <div class="content">{{content}}</div>
                </div>
            </div>
        </transition>

    </div>
</template>
<script>
export default {
    name: 'LeeToolTip',
    inheritAttrs: false,
    data() {
        return {
            isshow:false
        }
    },
    computed: {
        positionStyle() {
            if (this.position === 'left') {
                return 'left'
            } else if (this.position === 'right') {
                return 'right'

            } else if (this.position === 'bottom') {
                return 'bottom'

            } else {
                return 'top'
            }

        },
        transitionCss(){
            return 'left'
        }

    },
    methods: {
        enter() {
            this.isshow = true;
        },
        leave() {
            this.isshow = false;
        }
    },
    props: {
        color: {
            type: String,
            default: '#333' //默认default
        },
        content: {
            type: String,
            default: '' //默认default
        },
        position: {
            type: String,
            default: 'top' //默认default
        }
    }
};
</script>
<style scoped>
.lee-tooltip {
    position: relative;margin-right: 10px;display: inline-block;
}

.lee-tooltip-item {
    position: absolute;
    font-size: 12px;
}
.lee-tooltip-itemn{display: flex;}
.lee-tooltip-item.left{top: 50%;transform: translateY(-50%);right:100%;}
.lee-tooltip-item.left .jt{width:10px;}
.lee-tooltip-item.left i{top: 50%;transform: translateY(-50%);border-color: transparent transparent transparent var(--color)}

.lee-tooltip-item.right{top: 50%;transform: translateY(-50%);left: 100%}
.lee-tooltip-item.right .jt{width:10px;}
.lee-tooltip-item.right i{top: 50%;right:0px;transform: translateY(-50%);border-color: transparent var(--color) transparent transparent}

.lee-tooltip-item.top{top: -100%;height:100%;left: 50%;transform: translateX(-50%);}
.lee-tooltip-item.top .lee-tooltip-itemn{flex-wrap: wrap;}
.lee-tooltip-item.top .jt{height:10px;width: 100%}
.lee-tooltip-item.top i{border-color: var(--color) transparent transparent transparent;left: 50%;transform: translateX(-50%);}

.lee-tooltip-item.bottom{bottom: -100%;left: 50%;transform: translateX(-50%);height:100%;}
.lee-tooltip-item.bottom .lee-tooltip-itemn{flex-wrap: wrap;}
.lee-tooltip-item.bottom .jt{height:10px;width: 100%}
.lee-tooltip-item.bottom i{border-color: transparent transparent var(--color) transparent ;left: 50%;transform: translateX(-50%);}



.lee-tooltip-item i {
    position: absolute;   
    width: 0px;height: 0px;
    border-width: 5px;
    border-style: solid;display: block;
}
.lee-tooltip-item .jt{position: relative;}

.lee-tooltip-item .content {
    color: #fff;
    white-space: nowrap;
    padding: 6px 10px;
    border-radius: 3px;
    background: var(--color)
}
.left-leave-active,.left-enter-active {
    transition: all .5s;
}
.left-enter,.left-leave-to {
    opacity: 0;
    margin-right:-10px;
}
.left-enter-to,.left-leave {
    opacity: 1;
}
.right-leave-active,.right-enter-active {
    transition: all .5s;
}
.right-enter,.right-leave-to {
    opacity: 0;
    margin-left:-10px;
}
.right-enter-to,.right-leave {
    opacity: 1;
}
.top-leave-active,.top-enter-active {
    transition: all .5s;
}
.top-enter,.top-leave-to {
    opacity: 0;
    margin-top:-10px;
}
.top-enter-to,.top-leave {
    opacity: 1;
}
.bottom-leave-active,.bottom-enter-active {
    transition: all .5s;
}
.bottom-enter,.bottom-leave-to {
    opacity: 0;
    margin-bottom:-10px;
}
.bottom-enter-to,.bottom-leave {
    opacity: 1;
}

</style>