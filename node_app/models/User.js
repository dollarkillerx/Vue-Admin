const mongoose = require("mongoose");
const Schema = mongoose.Schema;

// Create Schema
const UserSchema = new Schema({
    name: {
        type: String,
        request: true
    },
    email: {
        type:String,
        request: true
    },
    password:{
        type:String,
        request: true
    },
    avatar:{
        type:String
    },
    data:{
        type:Date,
        default:Date.now
    }
});

module.exports = User = mongoose.model("users",UserSchema);