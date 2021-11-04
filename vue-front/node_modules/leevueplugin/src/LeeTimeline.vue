<template>
    <div class="lee-timeline">
        <div class="mark" v-if="limitnum!=0&&limitnum<totalnum"></div>
        <ul>
            <slot></slot>
        </ul>
        <div class="more" v-if="limitnum!=0&&limitnum<totalnum" @click="chuli"></div>
    </div>
</template>
<script>
export default {
    name: 'LeeTimeline',
    inheritAttrs: false,
    data() {
        return {
            limitnum: 0,
            totalnum:0
        };
    },
    provide() {
        return {
            timeline:this
        };
    },
    mounted(){
        this.totalnum=this.$slots.default.length
    },
    props: {
        reverse: {
            type: Boolean,
            default: false
        },
        value: {
            type: Number,
            default: 0
        },
        theme: {
            type: Number,
            default: 1
        },
        animate: {
            type: Boolean,
            default: false //默认default
        }
    },
    methods: {
        chuli() {
            this.limitnum = 0
            //this.$emit('input', 0)
        }
    },
    watch: {
        value: {
            immediate: true,
            handler(value) {
                this.limitnum = value
            }
        },
        reverse: {
            immediate: true,
            handler(value) {
                if (value) {
                    this.$slots.default = [...this.$slots.default].reverse();
                }
            }
        }
    }
};
</script>
<style scoped>
* {
    margin: 0px;
    padding: 0px;
    list-style: none
}

.lee-timeline {
    position: relative;
}

.lee-timeline .more {
    font-size: 14px;
    color: #3eaf7c;
    width: 30px;
    height: 30px;
    border-radius: 50%;
    box-shadow: 0 1px 3px rgba(0, 0, 0, .3);
    line-height: 30px;
    text-align: center;
    position: absolute;
    cursor: pointer;
    left: 50%;
    bottom: -6px;
    z-index: 44
}

.lee-timeline .more:before {
    content: "";
    width: 7px;
    height: 7px;
    border-color: #3eaf7c;
    position: absolute;
    border-style: solid;
    border-width: 0 0 2px 2px;
    transform: rotate(-45deg);
    left: 50%;
    margin-left: -4px;
    top: 50%;
    margin-top: -6px
}

.lee-timeline .mark {
    position: absolute;
    width: calc(100% - 28px);
    height: 150px;
    bottom: 0px;
    right: 0px;
    background: -webkit-linear-gradient(rgba(255,255,255,0), #fff);
    z-index: 33
}
</style>