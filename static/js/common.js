var URL_PATH    =   'http://39.108.50.204/index.php';       //  服务器地址
var API_SUCCESS =   200;                                    //  200
var API_ERROR   =   500;                                    //  500
var handleLogout;                                           //  退出登录
var nativeAjax;                                             //  原生Ajax请求


//***************************************************************************************//

/**
 * 公用退出登录
 * 
 */
handleLogout = function()
{
    element.$confirm('确定退出登录?', '退出登录',
    {
        confirmButtonText: '退出',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        window.location.href = 'auth/logout';
    }).catch(() => {
        //  取消
    });
}

//---------------------------------------------------------------------------------------//

/**
 * 原生Ajax请求
 * @param options
 * @param show_log [Boolean]
 */
nativeAjax = function(options, show_log=false)
{
    options = options || {};                                //  传入方式默认为对象
    options.type = (options.type || "GET").toUpperCase();   //  默认为GET请求
    options.dataType = options.dataType || 'json';          //  返回值类型默认为json
    options.async = options.async || true;                  //  默认为异步请求
    
    var params = nativeGetParams(options.data, show_log);   //  对需要传入的参数的处理
    var xhr;                                                //  创建一个ajax请求
    try {                                                   //  W3C标准
        xhr = new XMLHttpRequest();
    } catch (e) {                                           //  IE标准{ActiveXObject}
        try {
            xhr = new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e) {
            xhr = new ActiveXObject("Microsoft.XMLHTTP");
        }
    }
    
    if ( !xhr )
    {
        alert('你的浏览器不支持AJAX！请升级你的浏览器。');
        return false;
    }
    
    xhr.onreadystatechange = function()
    {
        if ( xhr.readyState == 4 )
        {
            var status = xhr.status;
            if ( status >= 200 && status < 300 ) {
                options.success && options.success(JSON.parse(xhr.responseText), xhr.responseXML);
            } else {
                options.fail && options.fail(status);
            }
        }
    };
    
    if ( options.type == 'GET' ) {
        xhr.open("GET",options.url + '?' + params ,options.async);
        xhr.send(null);
    } else if ( options.type == 'POST' ) {
        xhr.open('POST',options.url,options.async);                                 //  打开请求
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');  //  POST请求设置请求头
        // xhr.setRequestHeader("X-CSRF-TOKEN", document.querySelector('meta[name="csrf-token"]').getAttribute('content'));
        xhr.send(params);                                                           //  发送请求参数
    }
}

/**
 * 原生Ajax请求对象参数的处理
 * @param data
 * @returns {string}
 */
var nativeGetParams = function(data, show_log)
{
    var arr = [];
    for ( var param in data )
    {
        arr.push(encodeURIComponent(param) + '=' +encodeURIComponent(data[param]));
    }
    if ( show_log ) { console.log(arr); }
    arr.push(('randomNumber=' + Math.random()).replace('.'));
    if ( show_log ) { console.log(arr); }
    return arr.join('&');
}

//---------------------------------------------------------------------------------------//

