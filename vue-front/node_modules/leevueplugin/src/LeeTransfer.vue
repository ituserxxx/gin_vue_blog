<template>
    <div class="lee-transfer" :class="{filterable:filterable}">
        <div class="lee-transfer-panel">
            <div class="lee-transfer-header">
                <div class="header">
                    <div class="all">
                        <div @click="allleftchoose" class="checkbox" :class="{check:allleft}"><span></span>{{title1}}
                        </div>
                    </div>
                    <div class="js">{{leftxz}}/{{lefttj}}</div>
                </div>
            </div>
            <div class="lee-transfer-body">
                <div class="search" v-if="filterable"><i @click="keyword1=''" v-if="keyword1.length>0"></i><input
                        type="text" name="" v-model="keyword1" placeholder="搜索关键字"></div>
                <ul v-if="leftdata.length>0">
                    <li v-for="item in leftdata">
                        <div @click="item.choose=!item.choose" class="checkbox" :class="{check:item.choose==true}">
                            <span></span>{{item.label}}
                        </div>
                    </li>
                </ul>
                <p v-else>无数据</p>
            </div>
        </div>
        <div class="lee-transfer-button">
            <lee-button @click="toleft" type="leecolor" v-if="rightxz>0">&lt;</lee-button>
            <lee-button type="disabled" v-else>&lt;</lee-button>
            <lee-button @click="toright" type="leecolor" v-if="leftxz>0">&gt;</lee-button>
            <lee-button type="disabled" v-else>&gt;</lee-button>
        </div>
        <div class="lee-transfer-panel">
            <div class="lee-transfer-header">
                <div class="header">
                    <div class="all">
                        <div @click="allrightchoose" class="checkbox" :class="{check:allright}"><span></span>{{title2}}
                        </div>
                    </div>
                    <div class="js">{{rightxz}}/{{righttj}}</div>
                </div>
            </div>
            <div class="lee-transfer-body">
                <div class="search" v-if="filterable"><i @click="keyword2=''" v-if="keyword2.length>0"></i><input
                        type="text" name="" v-model="keyword2" placeholder="搜索关键字"></div>
                <ul v-if="rightdata.length>0">
                    <li v-for="item in rightdata">
                        <div @click="item.choose=!item.choose" class="checkbox" :class="{check:item.choose==true}">
                            <span></span>{{item.label}}
                        </div>
                    </li>
                </ul>
                <p v-else>无数据</p>
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        name: 'LeeTransfer',
        inheritAttrs: false,
        data() {
            return {
                keyword1: '',
                keyword2: ''
            }
        },
        mounted() {
            this.list.map((el) => {
                this.$set(el, 'choose', false) //choose表示选中状态
                this.$set(el, 'flag', 0) //flag=0表示在左边，flag=1表示在右边
            })
            for (var i = 0; i < this.value.length; i++) {
                var value = this.value[i].toString()
                var fi = this.list.filter(function (el) {
                    return el.key.toString().indexOf(value) > -1;
                });
                this.$set(fi[0], 'flag', 1)
            }

        },
        computed: {
            leftdata() {
                var _this = this
                var fi = this.list.filter(function (el) {
                    return el.flag == 0;
                });
                if (this.filterable) {
                    var fii = fi.filter(function (el) {
                        return el.label.toLowerCase().indexOf(_this.keyword1.toLowerCase()) > -1;
                    });
                    return fii
                } else {
                    return fi
                }

            },
            rightdata() {
                var _this = this
                var fi = this.list.filter(function (el) {
                    return el.flag == 1;
                });
                if (this.filterable) {
                    var fii = fi.filter(function (el) {
                        return el.label.toLowerCase().indexOf(_this.keyword2.toLowerCase()) > -1;
                    });
                    return fii
                } else {
                    return fi
                }
            },
            lefttj() {
                return this.leftdata.length
            },
            leftxz() {
                var fi = this.leftdata.filter(function (el) {
                    return el.choose == true
                });
                return fi.length
            },
            righttj() {
                return this.rightdata.length
            },
            rightxz() {
                var fi = this.rightdata.filter(function (el) {
                    return el.choose == true
                });
                return fi.length
            },
            allleft() {
                if (this.leftdata.length < 1) {
                    return false
                }
                var fii = this.leftdata.every(function (el) {
                    return el.choose == true;
                })
                return fii
            },
            allright() {
                if (this.rightdata.length < 1) {
                    return false
                }
                var fii = this.rightdata.every(function (el) {
                    return el.choose == true;
                })
                return fii
            },
        },
        methods: {
            allleftchoose() {
                var jg = this.leftdata.some(function (el) {
                    return !el.choose;
                })
                if (jg) {
                    this.leftdata.map(function (el) {
                        el.choose = true
                    })
                } else {
                    this.leftdata.map(function (el) {
                        el.choose = false
                    })
                }
            },
            allrightchoose() {
                var jg = this.rightdata.some(function (el) {
                    return !el.choose;
                })
                if (jg) {
                    this.rightdata.map(function (el) {
                        el.choose = true
                    })
                } else {
                    this.rightdata.map(function (el) {
                        el.choose = false
                    })
                }
            },
            toleft() {
                var _this = this
                var fi = this.rightdata.filter(function (el) {
                    return el.choose == true
                });
                fi.map(function (el) {
                    _this.$set(el, 'flag', 0)
                    _this.$set(el, 'choose', false)
                })
                this.changeV()
            },
            toright() {
                var _this = this
                var fi = this.leftdata.filter(function (el) {
                    return el.choose == true
                });
                fi.map(function (el) {
                    _this.$set(el, 'flag', 1)
                    _this.$set(el, 'choose', false)
                })
                this.changeV()
            },
            changeV() {
                let ll = []
                let fii = this.list.filter((el) => {
                    return el.flag == 1
                })
                fii.map((el) => {
                    ll.push(el.key)
                })
                this.$emit('input', ll)
            }
        },
        props: {
            list: {
                type: Array,
                default() {
                    return []
                }
            },
            value: {
                type: Array,
                default() {
                    return []
                }
            },
            title1: {
                type: String,
                default: ''
            },
            title2: {
                type: String,
                default: ''
            },
            filterable: {
                type: Boolean,
                default: false
            }

        }
    };
</script>
<style scoped>
    .lee-transfer-panel {
        border: 1px solid #ebeef5;
        border-radius: 4px;
        overflow: hidden;
        background: #fff;
        display: inline-block;
        vertical-align: middle;
        width: 200px;
        max-height: 100%;
        box-sizing: border-box;
        position: relative;
    }

    .lee-transfer-panel .lee-transfer-header {
        height: 40px;
        line-height: 40px;
        background: #f5f7fa;
        border-bottom: 1px solid #ebeef5;
        box-sizing: border-box;
        color: #000;
    }

    .lee-transfer-panel .lee-transfer-header .header {
        padding: 0 15px;
        height: 100%;
        display: flex;
        justify-content: space-between;
    }

    .lee-transfer-panel .checkbox {
        display: flex;
        align-items: center;
        cursor: pointer;
    }

    .lee-transfer-panel .checkbox:hover span {
        border: 1px solid #46bd87;
    }

    .lee-transfer-panel .checkbox:hover {
        color: #46bd87
    }

    .lee-transfer-panel .checkbox span {
        margin-right: 6px;
        border-radius: 3px;
        position: relative;
        border: 1px solid #dcdfe6;
        box-sizing: border-box;
        width: 14px;
        height: 14px;
        background-color: #fff;
        z-index: 1;
        transition: border-color .25s cubic-bezier(.71, -.46, .29, 1.46), background-color .25s cubic-bezier(.71, -.46, .29, 1.46);
    }

    .lee-transfer-panel .check span {
        background: #46bd87;
        border: 1px solid #46bd87;
    }

    .lee-transfer-panel .check span:after {
        content: '';
        height: 6px;
        width: 3px;
        left: 4px;
        position: absolute;
        top: 1px;
        border: 1px solid #fff;
        border-left: 0;
        border-top: 0;
        transform: rotate(45deg);
    }


    .lee-transfer-panel .lee-transfer-header .header .js {
        font-size: 12px;
        color: #aaa
    }

    .lee-transfer-panel .lee-transfer-body {
        height: 246px;
    }

    .lee-transfer-panel .lee-transfer-body p {
        text-align: center;
        margin: 0;
        padding: 12px 0;
    }

    .lee-transfer-panel .lee-transfer-body .check {
        color: #46bd87
    }

    .lee-transfer-panel .lee-transfer-body .search {
        margin: 12px 15px 0;
        position: relative;
    }

    .lee-transfer-panel .lee-transfer-body .search i {
        border-radius: 100%;
        width: 12px;
        height: 12px;
        box-sizing: border-box;
        border: solid 1px #dcdfe6;
        bacground: #fff;
        position: absolute;
        display: block;
        top: 50%;
        margin-top: -6px;
        right: 10px;
        cursor: pointer;
        transform: rotate(45deg);
        transition: all .6s
    }

    .lee-transfer-panel .lee-transfer-body .search i:before {
        content: '';
        width: 6px;
        height: 1px;
        position: absolute;
        background: #dcdfe6;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    .lee-transfer-panel .lee-transfer-body .search i:after {
        content: '';
        width: 1px;
        height: 6px;
        position: absolute;
        background: #dcdfe6;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }

    .lee-transfer-panel .lee-transfer-body .search i:hover {
        border-color: #46bd87;
        background: #46bd87
    }

    .lee-transfer-panel .lee-transfer-body .search i:hover:before,
    .lee-transfer-panel .lee-transfer-body .search i:hover:after {
        background: #fff
    }

    .lee-transfer-panel .lee-transfer-body .search input {
        width: 100%;
        height: 32px;
        border: 1px solid #dcdfe6;
        border-radius: 14px;
        outline: none;
        text-indent: 1em;
        box-sizing: border-box;
    }

    .lee-transfer-panel .lee-transfer-body .search input:focus {
        border-color: #46bd87
    }

    .lee-transfer-panel .lee-transfer-body ul,
    .lee-transfer-panel .lee-transfer-body li {
        list-style: none;
        margin: 0px;
        padding: 0px;
    }

    .lee-transfer-panel .lee-transfer-body ul {
        display: block;
        height: 246px;
        box-sizing: border-box;
        padding: 6px 15px;
        overflow: auto;
    }


    .lee-transfer-panel .lee-transfer-body ul li {
        height: 30px;
        line-height: 30px;
        /*padding: 0 15px;*/
        font-size: 14px;
        display: flex;
        align-items: center
    }

    .lee-transfer.filterable .lee-transfer-body {
        display: flex;
        flex-direction: column;
    }

    .lee-transfer-button {
        display: inline-block;
        vertical-align: middle;
        padding: 0 30px;
    }

    .lee-transfer-button .lee-button {
        display: block;
    }

    .lee-transfer-button .lee-button:nth-child(1) {
        margin-bottom: 15px;
    }
</style>