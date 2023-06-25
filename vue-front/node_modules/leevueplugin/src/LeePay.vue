<template>
    <div>
        <transition name="slideUp">
            <div class="PayChose" v-if="show">
                <div class="okdiv" @click="chuli" :style="{'background':config.bg}">关闭</div>
                <ul>
                    <li v-for="item in config.choose" @click="sel(item)" :style="{'color':tempay===item?config.bg:''}">{{item}}</li>
                </ul>
            </div>
        </transition>
        <transition name="fade">
            <div class="mark" @click="chuli" v-if="show"></div>
        </transition>
    </div>
</template>
<script>
export default {
    name: 'LeePay',
    data() {
        return {
            show: false,
            tempay: '',
        };
    },    
    props: {
        value: {
            type: Boolean,
            default: false
        },
        config: {
            type: Object,
            default() {
                return {
                    bg: '#409eff',
                    choose: ['微信支付', '支付宝', '花呗支付', '农业银行']
                }
            }
        }
    },
    methods: {
        sel(el) {
            this.tempay = el
        },
        chuli() {
            this.$emit('change',this.tempay)
            this.$emit('input',false)
        }
    },
    watch: {
        value: {
            immediate: true,
            handler(value) {
                this.show = value
            }
        }
    },
};
</script>
<style scoped>
* {
    margin: 0px;
    padding: 0px;
    list-style: none
}

.PayChose {
    position: fixed;
    width: 100%;
    height: auto;
    bottom: 0;
    left: 0;
    z-index: 66;
    background: #fff;
}


.PayChose li {
    line-height: 44px;
    border-top: solid 1px #ddd;
    padding: 0 15px;
    text-align: center;
    cursor: pointer;
}

.PayChose li.cur {}

.mark {
    position: fixed;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    background: rgba(0, 0, 0, .8);
    z-index: 36;
}


.okdiv {
    width: 15%;
    height: 25px;
    position: absolute;
    right: 0;
    top: -25px;
    line-height: 25px;
    text-align: center;
    color: #fff;
    font-size: 12px;
    cursor: pointer;
}




/*遮罩動畫*/

.fade-enter-active,
.fade-leave-active {
    transition: all .3s;
    visibility: visible;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
    visibility: hidden;
}


/*支付選擇動畫*/

.slideUp-enter-active,
.slideUp-leave-active {
    transition: transform .3s;
    transform: translate3d(0, 0, 0);
}

.slideUp-enter,
.slideUp-leave-to {
    transform: translate3d(0, 120%, 0);
}
</style>