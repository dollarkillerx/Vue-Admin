<template>
    <div class="fillContainer">
        <div>
            <el-form :inline="true" ref="add_data">
                <el-form-item class="btnRight">
                    <el-button type="primary" size="small" icon="view" @click="handleAdd()">添加</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="table_container">
        <el-table
                :data="tableData"
                max-height="450"
                border
                style="width: 100%">
            <el-table-column
                    prop="id"
                    label="Id"
                    align="center"
                    width="60">
            </el-table-column>
            <el-table-column
                    prop="name"
                    label="名称"
                    align="center"

            >
            </el-table-column>
            <el-table-column
                    prop="language"
                    label="语言"
                    align="center"

                    width="80">
            </el-table-column>
            <el-table-column
                    prop="describe"
                    label="备注"
                    align="center"

            >
            </el-table-column>
            <el-table-column
                    prop="create_time_str"
                    label="创建时间"
                    align="center"
            >
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ scope.row.create_time_str }}</span>
                </template>
            </el-table-column>
            <el-table-column
                    label="操作"
                    prop="operation"
                    align="center"
                    fixed="right"
                    width="160"
            >
                <template slot-scope="scope">
                    <el-button
                            type="warning"
                            size="mini"
                            icon="edit"
                            @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                    <el-button
                            size="mini"
                            type="danger"
                            icon="delete"
                            @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        </div>
        <Dialog :dialog="dialog"></Dialog>
    </div>
</template>

<script>
    import Dialog from "../components/Dialog";
    export default {
        name: "sortList",
        data() {
            return {
                tableData: [],
                dialog:{
                    show:false
                }
            };
        },
        created() {
            this.getProfile();
        },
        methods: {
            getProfile() {
                // 获取表格数据
                this.$axios
                    .get("/api/sort/page")
                    .then(res => {
                        this.tableData = res.data;
                        console.log(this.tableData)
                    })
                    .catch(err => {
                        console.log(err)
                    })
            },
            handleEdit(index,row) {
                console.log(row.id)
            },
            handleDelete(index,row) {
                console.log(row.id)
            },
            handleAdd() {
                this.dialog.show = true
            }
        },
        components: {
            Dialog
        }
    }
</script>

<style scoped>
    .fillContainer {
        width: 100%;
        height: 100%;
        padding: 16px;
        box-sizing: border-box;
    }
    .btnRight {
        float: right;
    }
    .pagination {
        text-align: right;
        margin-top: 10px;
    }
</style>