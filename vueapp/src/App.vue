<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  import base64 from 'base-64/base64'
  export default {
    name: "app",
    components: {},
    created(){
      if(localStorage.DKToken) {
        // 获取token
        let token = localStorage.DKToken;
        // 解析token
        let a = token.split(".")
        let  b = a[0]
        let jwt_string = base64.decode(b);
        let jwt = JSON.parse(jwt_string)

        // 存储数据
        this.$store.dispatch("setIsAutnenticated", !this.isEmpty(jwt));
        this.$store.dispatch("setUser", jwt);
      }
    },
    methods:{
      isEmpty(value) {
        return (
                value === undefined ||
                value === null ||
                (typeof value === "object" && Object.keys(value).length === 0) ||
                (typeof value === "string" && value.trim().length === 0)
        );
      }

    }
  }
</script>


<style>
html,
body,
#app {
  width: 100%;
  height: 100%;
}

</style>
