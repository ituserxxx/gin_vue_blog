<template>
  <div class="app-container">
  <el-upload class="avatar-uploader" 
      :action="qn_upload_address" 
      :data="qiniuData"
      :multiple = "true"
      :show-file-list="true" 
      :before-upload="beforeUpload"
      :on-success="handleSuccess">
      <img v-if="icon" :src="icon ? icon : imgUrl" class="avatar">
      <i v-else class="el-icon-plus avatar-uploader-icon"></i>
    </el-upload>

  

  </div> 
</template>

<script>

import {getUploadToken} from '@/api/upload'
export default {
   props:{
    icon:{type:String}
  },
  data() {
    return {
      imageUrl:[],  //图片数组
      qn_upload_address:"", //七牛云的上传地址(华北)
      domain:'', //七牛云的图片外链地址
       qiniuData:{
        key:'', //图片名字处理
        token:''
      },
    }
  },
   mounted(){
       getUploadToken().then((response) => {
          this.qiniuData.token = response.data.token
          this.domain = response.data.domain
          this.qn_upload_address = response.data.region
      })
  },
  methods: {
    //上传成功了
       handleSuccess(res) {
        let  imgurl = this.domain+"/"+res.key
       this.imageUrl.push(imgurl);
        console.log(this.imageUrl)
    },
    beforeUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2; 
      if (!isJPG) { 
        this.$message.error('上传头像图片只能是 JPG 格式!'); 
      } if (!isLt2M) { 
        this.$message.error('上传头像图片大小不能超过 2MB!'); 
      }
      this.qiniuData.key = `upload_pic_${file.name}`
      return isJPG && isLt2M;   
    } ,

  }
}
</script>
