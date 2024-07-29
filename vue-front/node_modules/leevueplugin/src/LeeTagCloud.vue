<template>
    <div class="lee-tag-cloud" :class="{close:closeable}">
        <transition name="fade">
            <div class="tip" v-if="exp" :style="position">该标签有<strong>{{setnum}}</strong>人设置</div>
        </transition>
        <ul>
            <li @click.stop="add(inx)" @mouseleave="close" @mouseenter="open(inx)" v-for="(item,inx) in tags" :style="{color:item.color,fontSize:item.size+'px'}">{{item.t}}<span class="closeico" v-if="closeable" @click.stop="remove(inx)">x</span></li>
        </ul>
    </div>
</template>
<script>
export default {
    name: 'LeeTagCloud',
    inheritAttrs: false,
    data() {
        return {
            tags: this.tag,
            exp: false,
            setnum: 0,
            px: 0,
            py: 0,
            pw: 0
        }
    },
    mounted() {
        var _this = this
        this.tags.map(function(el, ind) {
            var rnd = Math.floor((Math.random() + 1) * 14 + Math.random() * 2)
            _this.$set(el, 'size', rnd)
            _this.$set(el, 'color', _this.color())
        })
    },
    computed: {
        position() {
            return 'left:'+this.px+'px;top:'+(this.py-50)+'px;transform:translateX(-'+(160-this.pw)*0.5+'px)'
        }
    },
    methods: {
        open(xh) {
            this.px=event.target.offsetLeft
            this.py=event.target.offsetTop
            this.pw=event.target.offsetWidth
            this.exp = true
            var fi = this.tags.filter(function(el, index) {
                return index == xh;
            });
            this.setnum = fi[0].num
        },
        close() {
            this.exp = false
        },
        add(xh) {
            var _this = this
            var fi = this.tags.filter(function(el, index) {
                return index == xh;
            });
            fi.map(function(el) {
                _this.$set(el, 'num', el.num + 1)
            })
            this.setnum = fi[0].num
        },
        remove(xh) {
            this.tags.splice(xh, 1)
        },
        color() {
            var r = Math.floor(Math.random() * 255);
            var g = Math.floor(Math.random() * 255);
            var b = Math.floor(Math.random() * 255);
            var color = 'rgb(' + r + ',' + g + ',' + b + ' )';
            return color
        }
    },
    props: {
        tag: {
            type: Array,
            default () {
                return []
            }
        },
        closeable: {
            type: Boolean,
            default: false
        },
        num: {
            type: Boolean,
            default: false
        }
    }
};
</script>
<style scoped>
.lee-tag-cloud {
    position: relative;
}

.lee-tag-cloud .tip {
    position: absolute;
    width: 160px;
    height: 50px;color: #fff;
    background: #46bd87;
    box-shadow: 0 0 3px rgba(0, 0, 0, .3);
    border-radius: 12px;
    line-height: 50px;
    text-align: center;
    /*left: 50%;top: 50%;margin-left: -80px;margin-top: -25px;*/
    z-index: 33
}

.lee-tag-cloud .tip:after {
    content: '';
    position: absolute;
    width: 0px;
    height: 0px;
    border-width: 5px;
    border-style: solid;
    left: 50%;
    bottom: -5px;
    border-color: transparent transparent #46bd87 #46bd87;
    box-shadow: -1px 1px 1px rgba(0, 0, 0, .3);
    transform: rotate(-45deg);
}
.lee-tag-cloud .tip strong{color: #ff0}
.lee-tag-cloud ul {
    width: 40%;
    height: 100%;
    margin: 0 auto;
    list-style: none;
    padding: 0px;
    text-align: center;
}

.lee-tag-cloud ul li {
    padding: 5px 10px;
    display: inline-block;
    margin: 0;
}

.lee-tag-cloud.close ul li {
    transition: all .5s;
    cursor: pointer;
    position: relative;
}

.lee-tag-cloud.close ul li .closeico {
    position: absolute;
    font-size: 12px;
    width: 12px;
    height: 12px;
    display: block;
    line-height: 8px;
    cursor: pointer;
    border-radius: 100%;
    background: rgba(0, 0, 0, 0);
    color: #fff;
    transition: all .5s;
    top: 0;
    right: 0
}

.lee-tag-cloud.close ul li:hover .closeico {
    background: rgba(0, 0, 0, .6)
}

.lee-tag-cloud ul li:nth-child(1),
.lee-tag-cloud ul li:nth-child(5) {
    vertical-align: top;
}

.fade-enter-active,
.fade-leave-active {
    transition: all .6s;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
}


@media (max-width: 767px) {
    .lee-tag-cloud ul {
        width: 100%
    }
}
</style>