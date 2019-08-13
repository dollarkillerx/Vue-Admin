<template>
    <div class="dialog">
        <el-dialog
                title="添加分类"
                :visible.sync="dialog.show"
                :close-on-click-modal="false"
                :close-on-press-escape="false"
                :modal-append-to-body="false"
        >
        <div class="form">
            <el-form
                ref="form"
                :model="formData"
                :rules="form_urles"
                label-width="120px"
                style="margin: 10px;width: auto;"
            >
                <el-form-item label="名称:">
                    <el-input v-model="formData.name" placeholder="分类名称"></el-input>
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
                formData:{
                    name: '',
                    language: '',
                    describe: ''
                },
                format_type_list:[
                    "名称",
                    "语言",
                    "注释"
                ]
            }
        },
        props: {
            dialog:Object
        }
    }
</script>

<style scoped>

</style>