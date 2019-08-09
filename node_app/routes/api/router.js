// @login & register
const express = require("express");
const router = express.Router();

// $route GET api/users/test
// @desc 返回请求的JSON数据
// @access public
router.get("/test",(req,res)=>{
    res.json({msg:"login"})
});

module.exports = router;