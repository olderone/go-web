var element = new Vue(
{
    el: '#element',
    data:
    {
        default_active      :   'copy-copy',
        tableData           :   {},
        current_page        :   1
    },
    mounted: function()
    {
        var _this = this;
        nativeAjax(
        {
            url: "/index.php/index/copy/copy",                              //请求地址
            type: 'POST',                                                   //请求方式
            data: {},                                                       //请求参数
            dataType: "json",                                               // 返回值类型的设定
            async: false,                                                   //是否异步
            success: function (response, xml)
            {
                switch ( response.code )
                {
                    case 200:
                        //  copy
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
                            message : '服务器未知错误，请重试',
                            center  : true,
                            type    : 'warning'
                        });
                        break;
                }
            },
            fail: function (status)
            {
                _this.$message.error(
                {
                    message : '网络错误，请重试',
                    center  : true,
                    type    : 'error'
                });
            }
        });
    },
    methods:
    {
        //  分页
        handleCurrentChange: function(_page)
        {
            var _this = this;
            nativeAjax(
            {
                url: "/index.php/index/copy/copy",                              //请求地址
                type: 'POST',                                                   //请求方式
                data: {page:_page},                                              //请求参数
                dataType: "json",                                               // 返回值类型的设定
                async: false,                                                   //是否异步
                success: function (response, xml)
                {
                    switch ( response.code )
                    {
                        case 200:
                            //  copy
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
                                message : '服务器未知错误，请重试',
                                center  : true,
                                type    : 'warning'
                            });
                            break;
                    }
                },
                fail: function (status)
                {
                    _this.$message.error(
                    {
                        message : '网络错误，请重试',
                        center  : true,
                        type    : 'error'
                    });
                }
            });
        },
        
        //  判断是否为(JSON)空对象
        isEmptyObject: function (e)
        {
            var t;
            for (t in e)
                return !1;
            return !0
        }
    }
})