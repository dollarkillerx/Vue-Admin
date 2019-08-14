<template>
    <div class="dialog">
        <el-dialog
                :title="dialog.title"
                :visible.sync="dialog.show"
                :close-on-click-modal="false"
                :close-on-press-escape="false"
                :modal-append-to-body="false"
        >
        <div class="form">
            <el-form
                ref="form"
                :model="formData"
                :rules="form_rules"
                label-width="120px"
                style="margin: 10px;width: auto;"
            >
                <el-form-item prop="name" label="名称:">
                    <el-input type="name" v-model="formData.name" placeholder="分类名称"></el-input>
                </el-form-item>

                <el-form-item prop="language" label="语言:">
                    <el-input type="language" v-model="formData.language" placeholder="语言名称简写 比如西语: es"></el-input>
                </el-form-item>

                <el-form-item prop="describe" label="注释:">
                    <el-input type="describe" v-model="formData.describe" placeholder="注释"></el-input>
                </el-form-item>

                <el-form-item class="text_right">
                    <el-button @click="dialog.show = fa">
                        取消
                    </el-button>
                    <template v-if="dialog.type != 'edit'">
                        <el-button type="primary" @click="onSubmit('form')">
                            提交
                        </el-button>
                    </template>
                    <template v-else>
                        <el-button type="primary" @click="onEdit('form')">
                            修改
                        </el-button>
                    </template>
                </el-form-item>

            </el-form>
        </div>

        </el-dialog>
    </div>
</template>

<script>
    export default {
        name: "dialog",
        data(){
            return {

                format_type_list:[
                    "名称",
                    "语言",
                    "注释"
                ],
                form_rules: {
                    describe:[
                        {required:true,message:"收支描述不能为空",trigger:blur}
                    ],
                    language:[
                        {required:true,message:"语言不能为空",trigger:blur}
                    ],
                    name:[
                        {required:true,message:"名称不能为空",trigger:blur}
                    ]

                }
            }
        },
        props: {
            dialog:Object,
            formData:Object
        },
        methods:{
            onSubmit(form) {
                this.$refs[form].validate(valid => {
                    if(valid) {
                        // console.log(this.formData)
                        this.$axios.post("/api/sort/add",this.formData)
                            .then(res => {
                                // 添加成功
                                this.$message({
                                    message:'数据添加成功',
                                    type: 'success'
                                });
                                // 隐藏dialog
                                this.dialog.show = false;
                                // 告诉父组件添加成功刷新
                                this.$emit('update');
                            }).catch(err => {
                                // 添加失败
                                this.$message({
                                    message:'数据添加失败',
                                    type:'error'
                                })
                        })
                    }
                })
            },
            onEdit(form) {
                this.$refs[form].validate(valid => {
                    if(valid) {
                        // console.log(this.formData)
                        this.$axios.post("/api/sort/modify",this.formData)
                            .then(res => {
                                // 添加成功
                                this.$message({
                                    message:'数据添加成功',
                                    type: 'success'
                                });
                                // 隐藏dialog
                                this.dialog.show = false;
                                // 告诉父组件添加成功刷新
                                this.$emit('update');
                            }).catch(err => {
                            // 添加失败
                            this.$message({
                                message:'数据添加失败',
                                type:'error'
                            })
                        })
                    }
                })

            }
        },
    }
</script>

<style scoped>

</style>