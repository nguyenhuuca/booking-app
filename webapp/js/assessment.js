$.ajaxSetup({
    headers:{
        'Authorization': localStorage.getItem("jwt")
    }
})
var appConst = {
    baseUrl: "http://localhost:8080"
}


var shortType = {
    currentPrice: 'asc',
    currnentProdName: 'asc',
    currentBranchName: 'asc',
}


function ProductObj(id, name, branch, price) {
        this.id = id;
        this.name = name;
        this.branch = branch;
        this.price = price;
};


function bookingSearch() {
    $("#loginSpinner").show();
    $("#errMsg").text("");
    $("#errMsg").hide();
    var seachObj = {
        name: $("#prodName").val(),
        branch: $("#branchName").val(),
        price: $("#prodPrice").val()
    }

    $.ajax({
        url: appConst.baseUrl.concat("/products/filter"),
        type: "GET",
        data: seachObj,
        contentType: "application/json",
        dataType: "json"
    }).done(function(rs) {
        renderHTML(rs)
        $("#loginSpinner").hide();
    }).fail(function(err) {
        $("#errMsg").text(err.responseJSON.msg);
        $("#errMsg").show();
        $("#loginSpinner").hide();
    });
}

function shortPrice() {
    $("#errMsg").text("");
    $("#errMsg").hide();
    var seachObj = {
        name: 'price',
        type: shortType.currentPrice,
    }
    getSoftData(seachObj);

    if(shortType.currentPrice == 'asc') {
        shortType.currentPrice = 'desc'
    } else {
        shortType.currentPrice = 'asc'
    }

   
}

function shortName() {
    $("#errMsg").text("");
    $("#errMsg").hide();
    var seachObj = {
        name: 'name',
        type: shortType.currnentProdName,
    }
    getSoftData(seachObj);

    if(shortType.currnentProdName == 'asc') {
        shortType.currnentProdName = 'desc'
    } else {
        shortType.currnentProdName = 'asc'
    }

   
}

function shortBranch() {
    $("#errMsg").text("");
    $("#errMsg").hide();
    var seachObj = {
        name: 'branch',
        type: shortType.currentBranchName,
    }
    getSoftData(seachObj);

    if(shortType.currentBranchName == 'asc') {
        shortType.currentBranchName = 'desc'
    } else {
        shortType.currentBranchName = 'asc'
    }

   
}

function getSoftData(seachObj) {
    $.ajax({
        url: appConst.baseUrl.concat("/products/short"),
        type: "GET",
        data: seachObj,
        contentType: "application/json",
        dataType: "json"
    }).done(function(rs) {
        renderHTML(rs)
    }).fail(function(err) {
        $("#errMsg").text(err.responseJSON.msg);
        $("#errMsg").show();
    });

}


/**
 * When load page, need to init state for some element on page
 */
function initState() {
    $("#loginSpinner").hide();
    $("#messageInfo").hide();
    $("#errMsg").hide();

}

/**
 * Using to return fix template for list video, each video will have item with fix fomat
 * @returns html with paramter by format: {{param}}
 */
function loadTemplate() {
    return `
    <div class="row">
        <!-- 16:9 aspect ratio -->
        <div class="col-6">
        <div class="ratio ratio-16x9">
            <img src="assets/prod.png" alt = "product example">
        </div>
        </div>
        <div class="col-6">
            <div>
                <div style = "float:left;font-weight: 600;">Product:&nbsp;</div> 
                <div style="color:red; font-weight:bold;">{{prod_name}}</div>
            </div>
        
            <div>
                <div style = "float:left;font-weight: 600;">Branch:&nbsp;</div> <div>{{branch}}</div>          
            </div>
            <div>
                <div style = "float:left;font-weight: 600;" >Price:&nbsp;</div> <div>{{price}}</div>
            </div>
        </div>
    </div>
    </br>
    `;
}

/**
 * Using to replace some data by real data that get froms server
 * @param {*} prodObj hold all data need to binding to html
 * @param {*} templateHtml  the static html from loadTemplate
 * @returns  string html after replace all data
 */
function bindingDataWhenLoad(prodObj, templateHtml) {
    var stringHtml = templateHtml
    stringHtml = stringHtml.replace("{{prod_name}}", prodObj.name);
    stringHtml = stringHtml.replace("{{branch}}", prodObj.branch);
    stringHtml = stringHtml.replace("{{price}}", prodObj.price);
    stringHtml = stringHtml.replace("{{desc}}", prodObj.desc);
    return stringHtml;
}

/**
 * Using to get all link share when page is loaded
 */
function loadData() {
    $.ajax({
        url: appConst.baseUrl.concat("/products"),
        type: "GET",
        contentType: "application/json",
        dataType: "json"
    }).done(function(rs) {
        renderHTML(rs)
    }).fail(function(err) {
        $("#errMsg").text(err.responseJSON.mesage);
        $("#errMsg").show();
    });

}

function renderHTML(productInfoList) {
    $('#list-product').empty();
    if (productInfoList == null) {
        $("#errMsg").text("No product to show");
        $("#errMsg").show();
        return
    }
    var data = []
    console.log(productInfoList);
    productInfoList.forEach(item =>{
        const product = new ProductObj(item.id, item.name, item.branch, item.price);
        data.push(product);
    });
    
    data.forEach(item => {
        var templateHtml = loadTemplate();
        stringHtml = bindingDataWhenLoad(item, templateHtml); 
        $('#list-product').append(stringHtml);
    });
}


