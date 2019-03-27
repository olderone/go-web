var element = new Vue(
{
    el: '#element',
    data:
    {
        formLabelWidth: '120px',
        options: [],
        selectedOptions: ['0', '0', '0'],
        dialogFormVisible : false,
        dialogFormVisible1 : false,
        dialogFormVisible2 : false,
        selectedData: [],
        form: {
            api_name              :   ''
        },

        para: {
            api_name              :   '',
            p_key   : '',
            p_val   : ''
        },

        head: {
            api_name              :   '',
            h_key   : '',
            h_val   : ''
        }
    },
    
    mounted: function()
    {
       
    },
    methods:
    {
        handleRemove: function(file, fileList)
        {
            console.log(file, fileList);
        },
        
        handlePreview: function(file)
        {
            console.log(file);
        },
        
        handleExceed: function(files, fileList)
        {
            this.$message.warning(`当前限制选择 1 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`);
        },
        
        beforeRemove: function(file, fileList)
        {
            return this.$confirm(`确定移除 ${ file.name }？`);
        },

        add_api: function() {
            var _this = this;
            let _loading = _this.$loading(
            {
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading'
            });

            nativeAjax(
            {
                url: "/add_api",
                type: 'POST',
                data: {api_name:_this.form.api_name},
                dataType: "json",
                async: false,
                success: function (response, xml)
                {
                    _loading.close();
                    switch ( response.code )
                    {
                        case 200:
                            _this.tableData = response.data;
                            _this.current_page = _page;
                            break;
                        case 500:
                            _this.$message(
                            {
                                message : response.message,
                                center  : true,
                                type    : 'warning'
                            });
                            break;
                        default:
                            _this.$message(
                            {
                                message : '未知错误，请重试。',
                                center  : true,
                                type    : 'warning'
                            });
                            break;
                    }
                },
                fail: function (status)
                {
                    _loading.close();
                    _this.$message.error(
                    {
                        message : '网络错误，请重试。',
                        center  : true,
                        type    : 'error'
                    });
                }
            });
        },


        //  分页
        handleCurrentChange: function(_page)
        {
            var _this = this;
            
            let _loading = _this.$loading(
            {
                lock: true,
                text: 'Loading',
                spinner: 'el-icon-loading'
            });

            nativeAjax(
            {
                url: "/admin/getOrganizationList",
                type: 'POST',
                data: {page:_page,orgNo:_this.searchOrgNo},
                dataType: "json",
                async: false,
                success: function (response, xml)
                {
                    _loading.close();
                    switch ( response.code )
                    {
                        case 200:
                            _this.tableData = response.data;
                            _this.current_page = _page;
                            break;
                        case 500:
                            _this.$message(
                            {
                                message : response.message,
                                center  : true,
                                type    : 'warning'
                            });
                            break;
                        default:
                            _this.$message(
                            {
                                message : '未知错误，请重试。',
                                center  : true,
                                type    : 'warning'
                            });
                            break;
                    }
                },
                fail: function (status)
                {
                    _loading.close();
                    _this.$message.error(
                    {
                        message : '网络错误，请重试。',
                        center  : true,
                        type    : 'error'
                    });
                }
            });
        }


    }
});