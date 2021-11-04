<template>
    <div class="lee-collapse-item" :class="{active:collapse.value.indexOf(name)>-1}">
        <div class="itemtab" @click="chuli(name)">{{title}}<i></i></div>
        <transition name="leetran" @enter="enter" @after-enter="afterenter" @leave="leave" @after-leave="afterleave">
            <div class="itemcontentw" v-show="collapse.value.indexOf(name)>-1">
                <div class="itemcontent">
                    <slot></slot>
                </div>
            </div>
        </transition>
        <!-- :style="{display:collapse.value.indexOf(name)>-1?'block':'none'}" -->
    </div>
</template>
<script>
export default {
    name: 'LeeCollapseItem',
    inheritAttrs: false,
    data() {
        return {}
    },
    inject: ['collapse'],
    methods: {
        chuli(item) {
            if (this.collapse.accordion) {
                if (this.collapse.expname ==item) {
                   this.collapse.$emit('input','')
                } else {
                   this.collapse.$emit('input',item)
                }
            } else {
                if (this.collapse.expname.indexOf(item) < 0) {
                    this.collapse.expname.push(item)
                } else {
                    var index = this.collapse.expname.indexOf(item)
                    this.collapse.expname.splice(index, 1)
                }
            }
        },
        enter(el) {
            var h = el.scrollHeight
            el.style.setProperty('--h', h + 'px')
        },
        afterenter(el) {
            el.removeAttribute('style')
        },
        leave(el) {
            var h = el.scrollHeight
            el.style.setProperty('--h', h + 'px')
        },
        afterleave(el) {
            el.style.setProperty('--h', '')
        }
    },
    props: {
        title: {
            type: String,
            default: '' //默认default
        },
        name: {
            type: String,
            default: '' //默认default
        }
    }
};
</script>
<style scoped>
.lee-collapse-item .itemtab {
    display: flex;
    align-items: center;
    height: 48px;
    line-height: 48px;
    background-color: #fff;
    color: #303133;
    cursor: pointer;
    border-bottom: 1px solid #ebeef5;
    font-size: 13px;
    font-weight: 500;
    transition: border-bottom-color .3s;
    outline: none;
    justify-content: space-between;
}

.lee-collapse-item .itemtab i {
    border-width: 4px;
    border-color: #ddd #ddd transparent transparent;
    border-style: solid;
    transform: rotate(45deg);
    margin: -4px 4px 0 0;
    transition: all .4s;
}

.lee-collapse-item .itemcontentw {
    border-bottom: 1px solid #ebeef5;
    overflow: hidden;
}

.lee-collapse-item .itemcontent {
    padding-bottom: 25px;
    font-size: 13px;
    color: #303133;
    line-height: 1.769230769230769;
}

.lee-collapse-item.active .itemtab {
    border-bottom-color: transparent;
}

.lee-collapse-item.active .itemtab i {
    transform: rotate(135deg);
}

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