const express = require("express");
const mongoose = require("mongoose");
const app = express();

// import router
const router = require("./routes/api/router")

// Db config
const db = require("./config/keys").mongoUri;

mongoose.connect(
    db,
    { useNewUrlParser: true }
    )
    .then(()=>console.log("MongoDb Connected"))
    .catch(err=>console.log(err))

app.get("/",(req,res)=>{
    res.send("Hello World!")
})

// 使用routes
app.use("/",router)

const port = process.env.PORT || 5000;

app.listen(port,()=>{
    console.log(`Server is running on port ${port}`)
})

