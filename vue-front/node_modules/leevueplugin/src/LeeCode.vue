<template>
    <div class="leeallcode">
        <transition name="leetran" @enter="enter" @after-enter="afterenter" @leave="leave" @after-leave="afterleave">
            <div class="leecode" v-show="show">
                <slot></slot>
            </div>
        </transition>
        <div class="btnexp" @click="show=!show"><i :class="{up:show}"></i><span v-if="!show">显示代码</span><span v-else>隐藏代码</span></div>
    </div>
</template>
<script>
export default {
    name: 'LeeCode',
    data() {
        return {
            show: false
        };
    },
    methods: {
        enter(el) {
            var h = el.scrollHeight
            el.style.setProperty('--h', h + 'px')
        },
        afterenter(el){
            el.removeAttribute('style')
        },
        leave(el){
            var h = el.scrollHeight
            el.style.setProperty('--h', h + 'px')
        },
        afterleave(el){
            el.style.setProperty('--h', '')
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
.leeallcode{border-top: 1px solid #eaeefb;}
.leecode{will-change: height;overflow:hidden;padding:0 15px;background: #eee}
.btnexp{height:44px;display:flex;justify-content: center;align-items: center;cursor:pointer;border-top:solid 1px #eaecef;transition: all .3s;color: #ddd;margin-top: -1px;}
.btnexp i{
        border-width: 4px;
        border-color: #ddd #ddd transparent transparent;
        border-style: solid;transform: rotate(135deg);margin:-4px 4px 0 0;transition: all .3s;
}
.btnexp i.up{transform: rotate(-45deg);margin-top: 4px}

.btnexp span{opacity: 0;transition: all .3s;}
.btnexp:hover{background: #aaa}
.btnexp:hover i{border-color: #fff #fff transparent transparent;margin-right:14px;}
.btnexp:hover span{opacity: 1;color: #fff;}
.leetran-enter-active {
        animation: my-in .3s ease-in-out;
}

.leetran-leave-active {
        animation: my-out .3s ease-in-out;
    }

@keyframes my-in {
        0% {
            height: 0
        }
        100% {
            height: var(--h);
        }
}

@keyframes my-out {
        0% {
            height: var(--h);
        }
        100% {
            height: 0
        }
}
</style>