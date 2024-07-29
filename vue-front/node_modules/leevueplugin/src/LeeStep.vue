<template>
    <div class="lee-step" :style="'flex-basis: '+stepw+';color:'+step.color" v-if="step.direction==='vertical'">
        <div class="status">
            <div class="ico" :class="{active:step.active>=xh}"><span v-if="step.finish==='finish'||step.active<xh">{{xh}}</span><div v-else class="gou"></div></div>
            <div class="line" :class="{hide:xh>=total,active:step.active>=xh+1}"></div>
        </div>
        <div class="main">
            <div class="text" :class="{active:step.active>=xh}">{{title}}</div>
            <div class="ms" :class="{active:step.active>=xh}"><slot></slot></div>
        </div>
    </div>
    <div class="lee-step" :style="'flex-basis: '+stepw+';color:'+step.color" v-else-if="step.direction==='horizontal'&&step.center">
        <div class="status">
            <div class="ico" :class="{active:step.active>=xh}"><span v-if="step.finish==='finish'||step.active<xh">{{xh}}</span><div v-else class="gou"></div></div>
            <div class="line" :class="{hide:xh>=total,active:step.active>=xh+1}"></div>
        </div>
        <div class="main">
            <div class="text" :class="{active:step.active>=xh}">{{title}}</div>
            <div class="ms" :class="{active:step.active>=xh}"><slot></slot></div>
        </div>
    </div>
    <div class="lee-step" :style="'flex-basis: '+stepw+';color:'+step.color" v-else-if="step.direction==='horizontal'&&xh<total">
        <div class="status">
            <div class="ico" :class="{active:step.active>=xh}"><span v-if="step.finish==='finish'||step.active<xh">{{xh}}</span><div v-else class="gou"></div></div>
            <div class="line" :class="{hide:xh>=total,active:step.active>=xh+1}"></div>
        </div>
        <div class="main">
            <div class="text" :class="{active:step.active>=xh}">{{title}}</div>
            <div class="ms" :class="{active:step.active>=xh}"><slot></slot></div>
        </div>
    </div>
    <div class="lee-step" :style="'flex-basis: '+stepw+';max-width:'+maxw+';color:'+step.color" v-else-if="step.direction==='horizontal'&&xh>=total">
        <div class="status">
            <div class="ico" :class="{active:step.active>=xh}"><span v-if="step.finish==='finish'||step.active<xh">{{xh}}</span><div v-else class="gou"></div></div>
            <div class="line" :class="{hide:xh>=total,active:step.active>=xh+1}"></div>
        </div>
        <div class="main">
            <div class="text" :class="{active:step.active>=xh}">{{title}}</div>
            <div class="ms" :class="{active:step.active>=xh}"><slot></slot></div>
        </div>
    </div>
</template>
<script>
export default {
    name: 'LeeStep',
    inheritAttrs: false,
    inject: ['step'],
    data() {
        return {
            xh: '',
            total: ''
        }
    },
    mounted() {
        var children = this.$parent.$children
        this.xh = children.indexOf(this) + 1
        this.total = children.length

    },
    computed: {
        stepw() {
            if (!this.total) {
                return
            }
            var zhi = 100 / (this.total - 1)
            return zhi+'%'
        },
        maxw() {
            if (!this.total) {
                return
            }
            var zhi = 100 / this.total
            return zhi+'%'
        },
    },
    props: {
        title: {
            type: String,
            default: '' //默认default
        }
    }
};
</script>
<style scoped>

.lee-step {
    flex-shrink: 1
}
.normal .lee-step:last-child{
    flex-basis: auto!important;
    flex-shrink: 0
}

.lee-step .status {
    position: relative;
    width: 100%
}

.lee-step .status .line {
    background: #ccc;
    width: 100%;
    height: 2px;
    position: absolute;
    left: 0px;
    top: 50%;
    margin-top: -1px;
}

.lee-step .status .line.active {
    background: currentColor;
}

.lee-step .status .line.hide {
    display: none
}

.lee-step .status .ico {
    height: 24px;
    width: 24px;
    border: solid 2px #ccc;
    border-radius: 100%;
    background: #fff;
    z-index: 3;
    position: relative;
    text-align: center;
    line-height: 20px;
    font-weight: bold;
    color: #ccc
}

.lee-step .status .ico.active {
    color: currentcolor;
    border-color: currentcolor
}

.lee-step .status .ico.active .gou{position: relative;width: 100%;height: 100%;transform:translateX(-14%) rotate(-45deg);left: 50%;}
.lee-step .status .ico.active .gou:before{content: '';width: 2px;height: 20%;background:currentcolor;position: absolute;left: 0 }
.lee-step .status .ico.active .gou:after{content: '';height: 2px;width: 60%;background:currentcolor;position: absolute;left: 0;top: 20%; }

.lee-steps.fill  .lee-step .status .ico.active .gou:before,.lee-steps.fill  .lee-step .status .ico.active .gou:after{background:#fff;}
.lee-steps.fill  .lee-step .status .ico.active span{color:#fff;}

.lee-step .text {
    color: #444;
    line-height: 36px;
}
.lee-step .ms{font-size: 12px;padding-bottom:1.2em; color: #aaa;}

.lee-step .text.active,.lee-step .ms.active {
    color: currentcolor;
}

.lee-steps.vertical .lee-step .status{width: 24px;}
.lee-steps.vertical .lee-step .status .line{width: 2px;height: 100%;top: 1px;left: 50%;margin-left: -1px;}
.lee-steps.vertical .lee-step .main{flex-grow: 1;padding-left: 10px;}
.lee-steps.vertical .lee-step .main .text{line-height: 24px;padding-bottom: 8px}

.lee-steps.fill  .lee-step .status .ico.active{background: currentcolor;border-color: currentcolor;}
.lee-steps.center .main{text-align: center;}
.lee-steps.center .status .ico{left: 50%;margin-left: -12px;}
.lee-steps.center .status .line{left: 50%;}
.lee-steps.center .main .ms{padding-left: 20%;padding-right: 20%}

</style>