<template>
    <div class="el-input-number">
        <span role="button" class="el-input-number_decrease" :class="{'is-disabled':currentValue<=min?true:false}" @click="dec"><i>-</i></span>
        <span role="button" class="el-input-number_increase" :class="{'is-disabled':currentValue>=max?true:false}" @click="inc"><i>+</i></span>
        <div class="el-input">
            <input ref="input" type="text" v-bind="$attrs" class="el-input_inner" :value="currentValue" v-on="inputListeners">
        </div>
    </div>
</template>
<script>
export default {
    name: 'LeeInputNumber',
    inheritAttrs: false,
    data() {
        return {
            currentValue: 0
        };
    },
    mounted() {},
    props: {
        value: {
            type: Number,
            default: 1 //默认1
        },
        min: {
            type: Number,
            default: 1 //默认1
        },
        max: {
            type: Number,
            default: 1 //默认1
        }
    },
    methods: {
        dec() {
            this.currentValue--;
            if (this.currentValue < this.min) {
                this.currentValue = this.min
                return
            }
            this.$emit('input', this.currentValue);
            this.$emit('change', this.currentValue);

        },
        inc() {
            this.currentValue++;
            if (this.currentValue > this.max) {
                this.currentValue = this.max
                return
            }
            this.$emit('input', this.currentValue);
            this.$emit('change', this.currentValue);

        }
    },
    watch: {
        value: {
            immediate: true,
            handler(value) {
                var newVal = Number(value);
                if (isNaN(newVal)) {
                    newVal = this.currentValue
                }
                if (newVal >= this.max) newVal = this.max;
                if (newVal <= this.min) newVal = this.min;
                this.currentValue = newVal;
                this.$emit('input', newVal);

            }
        }
    },
    computed: {
        inputListeners() {
            var vm = this;
            return Object.assign({}, this.$listeners, {
                // 这里确保组件配合 `v-model` 的工作
                input: function(event) {
                    //var jg = Number(event.target.value)
                    //vm.$emit('input', jg);
                },
                change: function(event) {
                    var jg = Number(event.target.value)
                    if (isNaN(jg)) {
                        jg = vm.currentValue
                    }
                    if (jg >= vm.max) jg = vm.max;
                    if (jg <= vm.min) jg = vm.min;
                    vm.$refs.input.value = jg;
                    vm.$emit('change', jg)
                    vm.$emit('input', jg)
                }
            })
        }
    }
};
</script>
<style scoped>
.el-input-number {
    position: relative;
    display: inline-block;
    width: 180px;
    line-height: 38px;
}

.el-input-number_decrease,
.el-input-number_increase {
    position: absolute;
    z-index: 1;
    top: 1px;
    width: 40px;
    height: auto;
    text-align: center;
    background: #f5f7fa;
    color: #606266;
    cursor: pointer;
    font-size: 13px;
}

.el-input-number_decrease {
    left: 1px;
    border-radius: 4px 0 0 4px;
    border-right: 1px solid #dcdfe6;
}

.el-input-number_increase {
    right: 1px;
    border-radius: 0 4px 4px 0;
    border-left: 1px solid #dcdfe6;
}

.el-input {
    position: relative;
    font-size: 14px;
    display: inline-block;
    width: 100%;
}

.el-input_inner {
    -webkit-appearance: none;
    background-color: #fff;
    background-image: none;
    border-radius: 4px;
    border: 1px solid #dcdfe6;
    box-sizing: border-box;
    color: #606266;
    display: inline-block;
    font-size: inherit;
    height: 40px;
    line-height: 40px;
    outline: none;
    padding: 0 15px;
    transition: border-color .2s cubic-bezier(.645, .045, .355, 1);
    width: 100%;
}

.el-input-number .el-input_inner {
    -webkit-appearance: none;
    padding-left: 50px;
    padding-right: 50px;
    text-align: center;
}

.el-input-number_decrease.is-disabled,
.el-input-number_increase.is-disabled {
    color: #c0c4cc;
    cursor: not-allowed;
    /*pointer-events: none;*/
}
</style>