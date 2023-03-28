<template>
    <div class="passwordinput" v-if="isshow">
        <div class="inputarea">
            <div class="inputs">
                <div class="tit">
                    <div class="close" @click="close">x</div>请输入支付密码</div>
                <div class="inputk">
                    <div class="tip" :class="{ok:code==1}">{{tips}}</div>
                    <ul :class="myclass">
                        <li v-for="(item,inx) in ge" :class="{on:password[inx]}"><i></i></li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="keyarea">
            <ul>
                <li @click="ws">1</li>
                <li @click="ws">2</li>
                <li @click="ws">3</li>
                <li @click="ws">4</li>
                <li @click="ws">5</li>
                <li @click="ws">6</li>
                <li @click="ws">7</li>
                <li @click="ws">8</li>
                <li @click="ws">9</li>
                <li class="k"></li>
                <li @click="ws">0</li>
                <li class="k" @click="del"><i>x</i></li>
            </ul>
        </div>
        <div class="mark"></div>
    </div>
</template>
<script>
export default {
    name: 'LeePasswordInput',
    inheritAttrs: false,
    data() {
        return {
            password: [],
            isshow: false,
            tips: '',
            ge: [1, 1, 1, 1, 1, 1],
            code: ''
        };
    },
    props: {
        value: {
            type: Boolean,
            default: false
        },
        pass: {
            type: String,
            default: '123456'
        }
    },
    methods: {
        ws(e) {
            var n = e.target.innerText;
            if (this.password.length >= 6) {
                return
            } else {
                this.password.push(n)
            }

        },
        del() {
            this.tips = ""
            if (this.code == '') {
                this.password.pop()
                this.code = ''
            } else {
                this.password = []
                this.code = ''
            }

        },
        close() {
            this.$emit('input', false)
        }
    },
    watch: {
        value: {
            immediate: true,
            handler(value) {
                this.isshow=value
            }
        },
        password: {
            immediate: true,
            handler(value) {
                value = value.join("");
                if (value.length == 6) {
                    if (value != this.pass) {
                        this.tips = "密码错误"
                        this.code = '-1'
                    } else {
                        this.tips = "密码正确"
                        this.code = '1'
                    }
                    this.$emit('change', this.code)
                }
            }
        }
    },
    computed: {
        myclass(){
            if(this.code==-1){
                return 'err'
            }else if(this.code==1){
                return 'ok'
            }else{
                return ''
            }
        }

    }
};
</script>
<style scoped>
*{margin:0px;padding:0px;list-style: none}
.passwordinput {
    position: fixed;
    width: 100%;
    height: 100%;
    left: 0px;
    top: 0;
    z-index: 33
}

.passwordinput .inputarea {
    position: absolute;
    z-index: 3;
    width: 100%;
    height: 66%;
    left: 0px;
    top: 0
}

.passwordinput .inputarea .inputs {
    background: #eee;
    width: 96%;
    height: 32%;
    position: absolute;
    left: 50%;
    margin-left: -48%;
    top: 50%;
    margin-top: -20%;
    border-radius: 4px
}

.passwordinput .inputarea .inputs .tit {
    height: 30%;
    border-bottom: solid 1px #dedede;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #666
}

.passwordinput .inputarea .inputs .tit .close {
    position: absolute;
    left: 8px;
    top: 5px;
    font-size: 24px
}

.passwordinput .inputarea .inputs .inputk {
    height: 70%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
}

.passwordinput .inputarea .inputs .inputk .tip {
    font-size: 14px;
    width: 100%;
    position: absolute;
    top: 0;
    text-align: center;
    color: #f00;
    top: 3px
}

.passwordinput .inputarea .inputs .inputk .tip.ok {
    color: #67c23a
}

.passwordinput .inputarea .inputs .inputk ul {
    display: flex;
    height: 50%;
    background: #fff;
    border-radius: 4px;
    border: solid 1px #dedede;
    width: 90%;
}

.passwordinput .inputarea .inputs .inputk ul li {
    border-right: solid 1px #dedede;
    height: 100%;
    width: 16.67%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.passwordinput .inputarea .inputs .inputk ul li:last-child {
    border-right: none
}

.passwordinput .inputarea .inputs .inputk ul li i {
    width: 30%;
    border-radius: 100%;
    background: #333;
    height: 0;
    padding-bottom: 30%;
    display: none;
   
}


.passwordinput .inputarea .inputs .inputk ul li.on i {
    display: block
}

.passwordinput .inputarea .inputs .inputk ul.err li i {
    background: #f00
}

.passwordinput .inputarea .inputs .inputk ul.err {
    border-color: #f00
}

.passwordinput .inputarea .inputs .inputk ul.err li {
    border-color: #f00
}

.passwordinput .inputarea .inputs .inputk ul.ok li i {
    background: #67c23a
}

.passwordinput .inputarea .inputs .inputk ul.ok {
    border-color: #67c23a
}

.passwordinput .inputarea .inputs .inputk ul.ok li {
    border-color: #67c23a
}


.passwordinput .keyarea {
    position: absolute;
    left: 0;
    bottom: 0;
    font-size: 26px;
    width: 100%;
    height: 34%;
    z-index: 3;
    background: #fff;
    color: #333;
}

.passwordinput .keyarea ul {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    height: 100%;
}

.passwordinput .keyarea li {
    width: 33.33%;
    display: block;
    text-align: center;
    height: 25%;
    display: flex;
    align-items: center;
    justify-content: center;
    border: solid 1px #eee;
    box-sizing: border-box;
    border-right: none;
    border-bottom: none;
    cursor: pointer;
}

.passwordinput .keyarea li.k {
    background: #c2cbd1
}

.passwordinput .mark {
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, .6);
    z-index: 2;
    top: 0;
    left: 0
}
</style>