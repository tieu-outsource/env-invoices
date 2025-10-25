## Request
https://cskh.npc.com.vn/home/AccountNPC

## Response
<!DOCTYPE html>
<html lang="vi">
<head>
    <meta charset="UTF-8">
    <title>Trung tâm CSKH Điện lực miền Bắc</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1">
    <link rel="shortcut icon" href="/Content/BCA/imgs/favicon-19.png">
    <link rel="stylesheet" href="/Content/BCA/assets/font-awesome-4.7.0/css/font-awesome.min.css">
    <link rel="stylesheet" href="/Content/BCA/assets/uikit-3.0.3/css/uikit.min.css">
    <link rel="stylesheet" href="/Content/BCA/assets/jquery-ui-1.12.1/jquery-ui.min.css">
    <link rel="stylesheet" href="/Content/BCA/assets/jquery-ui-1.12.1/jquery-ui.theme.min.css">
    <link rel="stylesheet" href="/Content/BCA/assets/Semantic-UI-master/dist/semantic.min.css">
    <link rel="stylesheet" href="/Content/BCA/css/bootstrap.css">
    <link rel="stylesheet" href="/Content/BCA/css/animate.css">
    <link rel="stylesheet" href="/Content/BCA/css/font-icons.css" type="text/css">
    <link rel="stylesheet" href="/Content/BCA/css/style.css">
    <link rel="stylesheet" href="/Content/BCA/css/layout.css">
    <link rel="stylesheet" href="/Content/BCA/css/edit.css">
    <link rel="stylesheet" href="/Content/BCA/css/forgot_pass.css">
    <link href="/Content/BCA/css/StyleThuan.css" rel="stylesheet">

    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <!-- Bao gồm Bootstrap CSS và JS -->
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.bundle.min.js"></script>
    <!-- Bao gồm bootstrap-fileinput CSS và JS -->
    <link href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.4.9/css/fileinput.min.css" rel="stylesheet">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.4.9/js/fileinput.min.js"></script>

    <link rel="stylesheet" href="/Content/BCA/css/components/bs-filestyle.css" type="text/css" />



    

    <link href='https://fonts.googleapis.com/css?family=Inter' rel='stylesheet'>
    <!--JS-->



 

    
    
<script src="/Content/BCA/js/jquery.buttonLoader.js"></script>

    <script src="/Content/BCA/assets/uikit-3.0.3/js/uikit.min.js"></script>
    <script src="/Content/BCA/assets/uikit-3.0.3/js/uikit-icons.min.js"></script>
    <script src="/Content/BCA/assets/jquery-ui-1.12.1/jquery-ui.min.js"></script>
    <script src="/Content/BCA/assets/Semantic-UI-master/dist/semantic.min.js"></script>
    <script src="/Content/BCA/js/functions.js"></script>

    <script src="/Scripts/jquery.unobtrusive-ajax.min.js"></script>

    
<script>
        $(function () {
            let availableTags = [
                "Dịch vụ hỗ trợ",
                "Dịch vụ liên quan đến hợp đồng",
                "Dịch vụ cấp điện",
                "Lịch ghi chỉ số",
                "Hướng dẫn thủ tục",
                "Hướng dẫn thanh toán",
                "Tra cứu chỉ số",
                "Tra cứu tiền điện",
                "Giới thiệu",
                "Liên hệ",
            ];
            let availablelink = [
                "DichVuTrucTuyen/TuVanSuDungDichVuDienSPC",
                "DichVuTrucTuyenMember/ThayDoiCongSuatHoacCongToSPC",
                "DichVuTrucTuyen/DichVuCapDienMoi",
                "DichVuTTCSKH/DichVuTTCSKHSPC#8",
                "Home/HuongDanThuTucSPC",
                "Home/HuongDanThanhToanTienDienSPC",
                "DichVuTTCSKH/IndexSPC#0",
                "DichVuTTCSKH/IndexSPC#2",
                "Home/GioiThieuSPC",
                "Home/LienHeSPC",
            ];
            $("#tags")
            .on("keydown", function (event) {
                if (event.keyCode === $.ui.keyCode.TAB &&
                    $(this).autocomplete("instance").menu.active) {
                    event.preventDefault();
                }
                if (event.keyCode === 13) {
                    event.preventDefault();
                    return false;
                }
            })
            .autocomplete({
                minLength: 0,
                autoFocus: true,
                source: availableTags,
                select: function (event, ui) {
                    //lấy giá trị
                    var tmpLang = ui.item.value;
                    var tmpIndexTag = availableTags.indexOf(tmpLang);
                    var tmpLink = availablelink[tmpIndexTag];
                    window.location.href = 'https://www.cskh.evnspc.vn/' + tmpLink;
                    //gáng
                    //$('#frmTimKiem_IDTimKiem').val(tmpLang);
                    ////submit
                    //$('#frmTimKiem_Index').submit();
                }
            });
            //$("#tags").autocomplete({
            //    source: availableTags
            //});
        });
</script>


    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-2SD0TPBM8J"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag() { dataLayer.push(arguments); }
        gtag('js', new Date());

        gtag('config', 'G-2SD0TPBM8J');
    </script>
</head>
<body>
    <section id="root" class="uk-height-viewport uk-offcanvas-content">
        <div id="my-id" uk-offcanvas="overlay: true">
            <button class="uk-offcanvas-close border-0 bg-transparent" type="button"><img src="/Content/BCA/imgs/img_back_menu.png" alt=""></button>
            <div class="uk-offcanvas-bar uk-padding-remove bg-white">
                <ul class="uk-nav-default uk-nav-parent-icon mb-2" id="Menu_TrangChu" uk-nav>
                    <li class="uk-parent">
                        <a href="" style="text-transform: uppercase !important;">Giới thi&#234;̣u</a>
                        <ul class="uk-nav-sub" hidden>
                            
                            <li><a href="/Home/GioiThieuNPC" style="text-transform: none !important;">Trung tâm CSKH NPC</a></li>
                            <li><a href="/Home/ChinhSachBaoMatNPC" style="text-transform: none !important;">Ch&#237;nh s&#225;ch bảo mật</a></li>
                            <li><a href="/Home/CacTieuChuanDichVuNPC" style="text-transform: none !important;">Ti&#234;u chuẩn dịch vụ</a></li>
                            
                        </ul>
                    </li>
                    <li class="uk-parent">
                        <a href="#" style="text-transform: uppercase !important;">Hướng dẫn</a>
                        <ul class="uk-nav-sub">
                            <li><a href="/Home/CongCuTinhDienNangNPC">Hỗ trợ ước tính điện năng</a></li>
                            <li><a href="/Home/HuongDanThuTucNPC" style="text-transform: none !important;">Hướng dẫn</a></li>
                            <li><a href="/Home/HuongDanThanhToanTienDienNPC" style="text-transform: none !important;">Hướng dẫn thanh to&#225;n tiền điện</a></li>
                            <li><a href="/Home/TinhHoaDonTienDien" style="text-transform: none !important;">T&#237;nh h&#243;a đơn tiền điện</a></li>
                            <li><a href="/Home/DienMatTroiApMaiNPC" style="text-transform: none !important;">Điện mặt trời &#225;p m&#225;i</a></li>
                            <li><a href="/Home/TraCuuVanBanPhapLuatNPC" style="text-transform: none !important;">Tìm văn bản pháp luật</a></li>
                            
                            <li><a href="/Home/ThongTinHoiDapNPC" style="text-transform: none !important;">Tư vấn hỏi đáp</a></li>
                            
                            <li><a href="/Home/TriAnKhachHangNPC" style="text-transform: none !important;">Video hướng dẫn dịch vụ</a></li>
                        </ul>
                    </li>
                    <li class="uk-parent">
                        <a href="#" style="text-transform: uppercase !important;">Dịch vụ trực tuyến</a>
                        <ul class="uk-nav-sub">
                            <li><a href="/DichVuTrucTuyen/DichVuCapDienMoi" style="text-transform: none !important;">Dịch vụ cấp điện mới</a></li>
                            <li><a href="/DichVuTrucTuyenMember/DichVuLienQuanDenHopDong" style="text-transform: none !important;">Dịch vụ trong hợp đồng</a></li>
                            <li><a href="/DichVuTrucTuyen/DichVuHoTro" style="text-transform: none !important;">Trả lời tự động/ trực tuyến</a></li>
                            <li><a href="/DichVuTrucTuyen/ThanhToanTrucTuyenNPC_TTTT" style="text-transform: none !important;">Thanh to&#225;n trực tuyến</a></li>
                            
                        </ul>
                    </li>
                    <li class="uk-parent">
                        <a href="#" style="text-transform: uppercase !important;">Tra cứu</a>
                        <ul class="uk-nav-sub">
                            <li><a onclick="tracuu(16);" href="/DichVuTTCSKH/IndexNPC?index=16">Tra cứu hồ sơ</a></li>
                            <li><a onclick="tracuu(15);" href="/DichVuTTCSKH/IndexNPC?index=15">Sản lượng điện ngày</a></li>
                            <li><a onclick="tracuu(17);" href="/DichVuTTCSKH/IndexNPC?index=17">Tra cứu ảnh chỉ số</a></li>
                            <li><a onclick="tracuu(0);" href="/DichVuTTCSKH/IndexNPC?index=0">Tra cứu sản lượng</a></li>
                            <li><a onclick="tracuu(1);" href="/DichVuTTCSKH/IndexNPC?index=1" style="text-transform: none !important;">Tra cứu tiền điện</a></li>
                            <li><a onclick="tracuu(2);" href="/DichVuTTCSKH/IndexNPC?index=2" style="text-transform: none !important;">H&#243;a đơn tiền điện</a></li>
                            <li><a onclick="tracuu(3);" href="/DichVuTTCSKH/IndexNPC?index=3" style="text-transform: none !important;">Th&#244;ng tin nợ tiền điện</a></li>
                            <li><a onclick="tracuu(4);" href="/DichVuTTCSKH/IndexNPC?index=4" style="text-transform: none !important;">Th&#244;ng tin thanh to&#225;n tiền điện</a></li>
                            <li><a onclick="tracuu(7);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=7" style="text-transform: none !important;">Lịch ngừng giảm cung cấp điện</a></li>
                            <li><a onclick="tracuu(8);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=8" style="text-transform: none !important;">Lịch ghi chỉ số</a></li>
                            <li><a onclick="tracuu(9);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=9" style="text-transform: none !important;">Điểm thu tiền điện</a></li>
                            <li><a onclick="tracuu(10);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=10" style="text-transform: none !important;">Th&#244;ng tin gi&#225; b&#225;n điện</a></li>
                            <li><a onclick="tracuu(14);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=14" style="text-transform: none !important;">Theo dõi tiến độ dịch vụ</a></li>
                            <li><a onclick="tracuu(5);" href="/DichVuTTCSKH/IndexNPC?index=5" style="text-transform: none !important;">Theo d&#245;i tiến độ c&#225;c dịch vụ kh&#225;c</a></li>
                            <li><a onclick="tracuu(6);" href="/DichVuTTCSKH/IndexNPC?index=6" style="text-transform: none !important;">Sản lượng điện mặt trời</a></li>
                            
                            <li><a onclick="tracuu(13);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=13" style="text-transform: none !important;">Tìm kiếm phòng giao dịch</a></li>
                            <li><a href="/Congbothongtin.html" style="text-transform: none !important;">Thông tin chất lượng dịch vụ</a></li>
                        </ul>
                    </li>

                    <li class="uk-parent">
                        <a href="#" style="text-transform: uppercase !important;">Tin tức</a>
                        <ul class="uk-nav-sub">
                            <li><a href="/Home/TinMoiEVNNPCCC" style="text-transform: none !important;">Tin ngành điện</a></li>
                            <li><a href="/Home/TinMoiEVNNPC" style="text-transform: none !important;">Tin nội bộ</a></li>
                            
                        </ul>
                    </li>
                    <li class="uk-parent">
                        <a href="#" style="text-transform: uppercase !important;">Li&#234;n hệ</a>
                        <ul class="uk-nav-sub">
                            <li><a href="/Home/LienHeNPC" style="text-transform: none !important;">Trung tâm CSKH</a></li>
                            <li><a href="/TinTuc/ChiTietTinNPC?MA_TBAI=238" style="text-transform: none !important;">Đến các đơn vị</a></li>
                        </ul>
                    </li>

                    
                </ul>
            </div>
        </div>

         

<header id="header-uk">
    <div class="top">
        <div class="uk-container">
            <div class="uk-flex uk-flex-middle uk-flex-between">
                <div class="support">
                    <i class="fa fa-phone" aria-hidden="true"></i>
                    Hỗ trợ trực tuyến 24/7:
                    <span>19006769</span>
                </div>
                

            </div>
        </div>
    </div>
    <div class="bottom bg-white uk-box-shadow-small uk-position-z-index" style="z-index: 1000;" uk-sticky>
        <div class="uk-container" id="TopMenu">
            <nav class="uk-navbar-container uk-navbar-transparent" uk-navbar>
                <div class="uk-navbar-left">
                    <a style="" href="/Home/IndexNPC"><img src="/Content/imageNPC/Logo/logo.svg" alt=""></a>
                </div>
                <div class="uk-navbar-right nav-overlay uk-visible@m">
                    <ul class="uk-navbar-nav" id='Menu_TrangChu'>
                        <li>
                            <a href="#">Giới thi&#234;̣u</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    
                                    <li><a href="/Home/GioiThieuNPC">Trung tâm CSKH NPC</a></li>
                                    <li><a href="/Home/ChinhSachBaoMatNPC">Ch&#237;nh s&#225;ch bảo mật</a></li>
                                    <li><a href="/Home/CacTieuChuanDichVuNPC">Ti&#234;u chuẩn dịch vụ</a></li>
                                    
                                </ul>
                            </div>
                        </li>
                        <li>
                            <a href="#">Hướng dẫn</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    <li><a href="/Home/CongCuTinhDienNangNPC">Hỗ trợ ước tính điện năng</a></li>
                                    <li><a href="/Home/HuongDanThuTucNPC">Hướng dẫn thủ tục MBĐ</a></li>
                                    <li><a href="/home/HuongDanThanhToanTienDienNPC">Hướng dẫn thanh to&#225;n tiền điện </a></li>
                                    <li><a href="/Home/TinhHoaDonTienDien">T&#237;nh h&#243;a đơn tiền điện</a></li>
                                    <li><a href="/Home/DienMatTroiApMaiNPC">Điện mặt trời &#225;p m&#225;i</a></li>
                                    <li><a href="/Home/TraCuuVanBanPhapLuatNPC">Tìm văn bản pháp luật</a></li>
                                    
                                    <li><a href="/Home/ThongTinHoiDapNPC">Tư vấn hỏi đáp</a></li>
                                    
                                    <li><a href="/Home/TriAnKhachHangNPC" style="text-transform: none !important;">Video hướng dẫn dịch vụ</a></li>
                                    <li><a href="/Tracuu/huongdanDVC" >HDSD các dịch vụ trên cổng DVCQG</a></li>
                                </ul>
                            </div>
                        </li>
                        <li>
                            <a href="#">Dịch vụ trực tuyến</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    <li><a href="/DichVuTrucTuyen/DichVuCapDienMoi">Dịch vụ cấp điện mới</a></li>
                                    <li><a href="/DichVuTrucTuyenMember/DichVuLienQuanDenHopDong">Dịch vụ trong hợp đồng</a></li>
                                    
                                    <li><a href="/DichVuTrucTuyen/DichVuHoTro">Trả lời tự động/ trực tuyến</a></li>
                                    <li><a href="/DichVuTrucTuyen/ThanhToanTrucTuyenNPC_TTTT">Thanh to&#225;n trực tuyến</a></li>
                                    <li><a href="/DichVuTrucTuyen/NangLuongMatTroi">Đăng ký bán điện MTMN</a></li>
                                </ul>
                            </div>
                        </li>
                        <li>
                            <a href="#">Tra cứu</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    <li><a onclick="tracuu(16);" href="/DichVuTTCSKH/IndexNPC?index=16">Tra cứu hồ sơ</a></li>
                                    <li><a onclick="tracuu(15);" href="/DichVuTTCSKH/IndexNPC?index=15">Sản lượng điện ngày</a></li>
                                    <li><a onclick="tracuu(17);" href="/DichVuTTCSKH/IndexNPC?index=17">Tra cứu ảnh chỉ số</a></li>
                                    <li><a onclick="tracuu(0);" href="/DichVuTTCSKH/IndexNPC?index=0">Tra cứu sản lượng</a></li>
                                    <li><a onclick="tracuu(1);" href="/DichVuTTCSKH/IndexNPC?index=1">Tra cứu tiền điện</a></li>
                                    <li><a onclick="tracuu(2);" href="/DichVuTTCSKH/IndexNPC?index=2">H&#243;a đơn tiền điện</a></li>
                                    <li><a onclick="tracuu(3);" href="/DichVuTTCSKH/IndexNPC?index=3">Th&#244;ng tin nợ tiền điện</a></li>
                                    <li><a onclick="tracuu(4);" href="/DichVuTTCSKH/IndexNPC?index=4">Lịch sử thanh toán tiền điện</a></li>
                                    <li><a onclick="tracuu(7);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=7">Lịch ngừng giảm cung cấp điện</a></li>
                                    <li><a onclick="tracuu(8);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=8">Lịch ghi chỉ số</a></li>
                                    <li><a onclick="tracuu(9);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=9">Điểm thu tiền điện</a></li>
                                    <li><a onclick="tracuu(10);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=10">Th&#244;ng tin gi&#225; b&#225;n điện</a></li>
                                    
                                    <li><a onclick="tracuu(14);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=14">Theo d&#245;i tiến độ cấp điện</a></li>
                                    <li><a onclick="tracuu(5);" href="/DichVuTTCSKH/IndexNPC?index=5">Theo dõi tiến độ dịch vụ</a></li>
                                    <li><a onclick="tracuu(6);" href="/DichVuTTCSKH/IndexNPC?index=6">Sản lượng điện mặt trời</a></li>
                                    
                                    <li><a onclick="tracuu(13);" href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=13">Tìm kiếm phòng giao dịch</a></li>
									<li><a  href="/Congbothongtin.html" style="text-transform: none !important;">Thông tin chất lượng dịch vụ</a></li>
                                </ul>
                            </div>
                        </li>

                        <li>
                            <a href="#">Tin tức</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    
                                    <li><a href="/Home/TinMoiEVNNPCCC" style="text-transform: none !important;">Tin ngành điện</a></li>
                                    <li><a href="/Home/TinMoiEVNNPC" style="text-transform: none !important;">Tin nội bộ</a></li>
                                    
                                    
                                </ul>
                            </div>
                        </li>
                        <li>
                            <a href="#">Li&#234;n hệ</a>
                            <div class="uk-navbar-dropdown uk-margin-remove">
                                <ul class="uk-nav uk-navbar-dropdown-nav">
                                    <li><a href="/Home/LienHeNPC" style="text-transform: none !important;">Trung tâm CSKH</a></li>
                                    <li><a href="/TinTuc/ChiTietTinNPC?MA_TBAI=238" style="text-transform: none !important;">Đến các đơn vị</a></li>
                                </ul>
                            </div>
                        </li>
                        

                    </ul>

                </div>
                <div class="uk-navbar-right nav-overlay tool">
                    <ul class="uk-navbar-nav">
                        <li>
                            <a class="uk-navbar-toggle" href="#" uk-search-icon ></a>
                            <div class="uk-navbar-dropdown search1"
                                 uk-drop="mode: click; cls-drop: uk-navbar-dropdown; boundary: !nav">
                                <div class="uk-grid-small uk-flex-middle" uk-grid>
                                    <div class="uk-width-expand">
                                        
                                        <form action="/home/TimKiemSPC" method="post"
                                              class="uk-search uk-search-default uk-width-1-1"
                                              id="frmTimKiem_Index">
                                            <span uk-search-icon></span>
                                            <input id="tags" class="uk-search-input border-0" type="search"
                                                   placeholder="T&#236;m kiếm">
                                            <input type="hidden" id="frmTimKiem_IDTimKiem" name="ChuoiTimKiem" />
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </li>
						<li>
							<div class="mdl-fab-expandable--child center" style="margin-top:25px;text-aline:center " >
								<a style="width: 20px; height:20px; min-width:22px" id="" onclick="showmodal_call()" 
								class="mdl-button mdl-js-button mdl-button--fab mdl-button--mini-fab mdl-button--colored" data-upgraded=",MaterialButton">

									<i class="material-icons" style="font-size: 14px; line-height:24px">phone</i>
								</a>
							</div>
						</li>
                        <li >
                            
<a href="/home/AccountNPC" title="Đăng nhập/đăng ký"><span uk-icon="user"></span></a>


                        </li>
                        <li class="uk-hidden@m"><a href="#my-id" class="pr-0 uk-navbar-toggle" uk-toggle><span uk-icon="menu"></span></a></li>
                    </ul>
                </div>
            </nav>
        </div>
    </div>
</header>
<form method="post" id="frmLanguages" action="/Languages/ChangeLanguage" hidden>
    <input type="hidden" id="frmLang_KyHieu" name="lang" />
</form>
<script type="text/javascript">
	function showmodal_call(){
		$("#exampleModalclick").show();
		
	}
	function hidemodal_call(){
		$("#exampleModalclick").hide();
		
	}
    function frmLanguages_Onclick_Vi() {
        var tmpLang = "Vi";

        $('#frmLang_KyHieu').val(tmpLang);
        $('#frmLanguages').submit();
    }
    function frmLanguages_Onclick_En() {
        var tmpLang = "En";

        $('#frmLang_KyHieu').val(tmpLang);
        $('#frmLanguages').submit();
    }
</script>

  


<div class="login-container">
    <div class="block_g uk-section-small uk-background-cover" style="background-image: url('../Content/BCA/imgs/bg/2-14.jpg')">
        <div class="uk-container">
            <h1 class="title uk-text-uppercase uk-h2 uk-text-white">T&#224;i khoản</h1>
            <ul class="uk-breadcrumb">
                <li><a href="#">Trang chủ</a></li>
                <li><span>T&#224;i khoản</span></li>
            </ul>
        </div>
    </div>
    <div class="uk-section bg-white" id="login-form-content">
        <!-- Phần chọn tài khoản đăng nhập -->

        <div id="login-selection" class="login-container text-center">
            <div class="login-wrapper">
                <h3 class="title_c uk-text-uppercase uk-margin-remove">Đăng nhập</h3>
                <br />
                <p style="color: #000f9f;">Chọn loại tài khoản bạn muốn sử dụng đăng nhập</p>

                <div class="d-flex justify-content-center gap-3">
                    <button class="btn btn-outline-primary custom-btn" id="btnTaiKhoan">
                        <!--<img src="../Content/BCA/imgs/EVN.svg" alt="Doanh nghiệp" class="me-2 btn-icon" style="height: 50px;">-->
                        <img src="../FileUploads/EVN-hover.svg" alt="Doanh nghiệp" class="me-2 btn-icon-hover" style="height: 50px;">
                        Mã Khách hàng sử dụng điện
                    </button>

                    <button class="btn btn-outline-success custom-btn" onclick="send_data_DVC()">
                        <!--<img src="../Content/BCA/imgs/VNID.svg" alt="Công dân" class="me-2 btn-icon" style="height: 50px;">-->
                        <img src="../FileUploads/VNID-hover.svg" alt="Công dân" class="me-2 btn-icon-hover" style="height: 50px;" >
                        Tài khoản Định danh điện tử
                    </button>
                </div>
            </div>
        </div>


        <div class="uk-container" style="display: none;" id="login-form-container">

            <div class="uk-grid uk-flex-center">
                <div class="uk-width-2-5@s">
                    <div class="tab_acc uk-subnav-pill uk-child-width-expand uk-margin-remove" uk-switcher>
                        <h2 class="title_c uk-text-uppercase uk-margin-remove center">ĐĂNG NHẬP</h2>
                    </div>
                    <div class="uk-switcher uk-margin">

                        <div>
                            <form id="login-form" name="login-form" class="form_acc" action="/Account/Login" method="post">
                                <input name="__RequestVerificationToken" type="hidden" value="JSQ242oFqVOIlU48ISkf1S2hMVpfXus1JFtJKPXwHJfEkPGUuYYns1ReeHvFnJOksoaL3FeTkUMMcRcrN4h6Wq3K6noGxTsuQk09rUkvVIU1" />
                                <div class="uk-margin-small">
                                    <div class="uk-position-relative">
                                        <span class="uk-form-icon" uk-icon="icon: user"></span>
                                        <input name="Username"
                                               onkeyup="frmDangNhap_OnkeyUp_MKH_DN();"
                                               id="frmDangKy_TenDangNhap_DN" class="uk-input" type="text" placeholder="T&#234;n đăng nhập">
                                    </div>
                                </div>

                                <div class="uk-margin-small" >
                                    <div class="uk-position-relative">
                                        <span class="uk-form-icon" uk-icon="icon: unlock"></span>
                                        <input name="Password" onkeyup="frmDangNhap_OnkeyUp_MKH()" id="frmDangKy_MatKhau_DN" class="uk-input" type="password" placeholder="Mật Khẩu">
                                        <span toggle="#frmDangKy_MatKhau_DN" class="fa fa-eye toggle-password uk-position-center-right unlock_pass" style="right: 15px;"></span>
                                    </div>

                                </div>

                                <div class="uk-margin-small uk-text-center">
                                    <div>
                                        
<script type="text/javascript">
$(function () {$('#cbb332284b824860b2a806de5b2da4ea').show();});
function ______43f0b906e159464fa99db8993423cdaa________() { $('#cbb332284b824860b2a806de5b2da4ea').hide(); $.post("/DefaultCaptcha/Refresh", { t: $('#CaptchaDeText').val() }, function(){$('#cbb332284b824860b2a806de5b2da4ea').show();}); return false; }</script> 
<br/>
<img id="CaptchaImage" src="/DefaultCaptcha/Generate?t=0ed7c9a74f0841849a46da5a45dc2f53"/><input id="CaptchaDeText" name="CaptchaDeText" type="hidden" value="0ed7c9a74f0841849a46da5a45dc2f53" /> <br/><a href="#CaptchaImage" id="cbb332284b824860b2a806de5b2da4ea" onclick="______43f0b906e159464fa99db8993423cdaa________()" style="display:none;">L&#224;m mới</a><br/>Nhập hình ảnh kiểm tra<br/><input autocomplete="off" autocorrect="off" data-val="true" data-val-required="Thông tin bắt buộc" id="CaptchaInputText" name="CaptchaInputText" type="text" value="" /><br/><span class="field-validation-valid" data-valmsg-for="CaptchaInputText" data-valmsg-replace="true"></span>
                                    </div>
                                </div>
                                <div class="uk-margin-small">
                                    <p style="color:red"></p>
                                </div>
                                <button class="btn btn-primary w-100 mt-3 btn-send-contact">Đăng nhập</button>
                                
                                <input id="previousLink" name="previousLink" type="hidden" value="" />
                            </form>

                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>



<script>
        $(document).ready(function () {
            $('#btnTaiKhoan').click(function () {
                $('#login-selection').hide();
                $('#login-form-container').show(); // Hiện form đăng nhập cũ
            });

            
        });

        $(document).ready(function () {
            // Khi nhấn nút "Đăng nhập tài khoản", hiển thị form login
            $('#btn-show-login').click(function () {
                $('#login-form-content').fadeIn();  // Hiện form
                $('#btn-show-login').hide(); // Ẩn nút đăng nhập để tránh hiển thị lại
            });
        });
</script>

<style>
    .custom-btn {
        width: 250px;
        height: 160px;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
        border: 2px solid #000;
        margin: 10px;
    }

    .btn-outline-primary {
        color: #000f9f;
        border-color: #000f9f;
        background-color: transparent;
        transition: all 0.3s ease-in-out;
    }

        .btn-outline-primary:hover {
            background-color: #000f9f;
            color: #fff;
            border-color: #000f9f;
        }

    .btn-outline-success {
        color: #1a5613;
        border-color: #1a5613;
        background-color: transparent;
        transition: all 0.3s ease-in-out;
    }

        .btn-outline-success:hover {
            background-color: #1a5613;
            color: #fff;
            border-color: #ffcb05;
        }

    .login-wrapper {
        border: 2px solid #000f9f;
        ; /* Viền đen */
        border-radius: 10px; /* Bo góc */
        padding: 75px; /* Khoảng cách giữa nội dung và viền */
        width: fit-content; /* Chỉ bao quanh nội dung */
        margin: auto; /* Căn giữa màn hình */
        background: linear-gradient(135deg, #fff, #f3f3f3);
        ; /* Nền trắng */
        box-shadow: 0px 0px 15px rgba(0, 0, 0, 0.3);
    }


    .input-group-text {
        background: #e9ecef;
        border: none;
    }

    .form-control {
        height: 45px;
    }

    .btn-primary {
        background-color: #000f9f;
        color: #fff;
        border: none;
    }

        .btn-primary:hover {
            background-color: #000c7a;
        }

    .text-primary {
        color: #000f9f !important;
    }

    .toggle-password {
        background: none;
        border: none;
        cursor: pointer;
        padding: 5px;
    }

    /* form login */
    /* Căn giữa form trong trang */
    .form_acc {
        max-width: 400px;
        width: 100%;
        min-width: 280px;
        margin: auto;
        padding: 20px;
        background: #fff;
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

        /* Tiêu đề "Đăng Nhập" */
        .form_acc h2 {
            text-align: center;
            font-size: 22px;
            font-weight: bold;
            margin-bottom: 15px;
            color: #000f9f;
        }

    /* Input fields */
    .uk-input {
        width: 100%;
        height: 40px;
        border-radius: 5px;
        border: 1px solid #ccc;
        padding: 10px;
        font-size: 14px;
    }

    /* Mật khẩu có icon hiển thị */
    .uk-position-relative {
        display: flex;
        align-items: center;
        position: relative;
    }

    .uk-form-icon {
        position: absolute;
        left: 10px;
        color: #999;
    }

    /* Nút đăng nhập */
    .btn-send-contact {
        width: 100%;
        height: 40px;
        background-color: #000f9f;
        color: white;
        font-size: 16px;
        font-weight: bold;
        border: none;
        border-radius: 5px;
        margin-top: 10px;
    }

        .btn-send-contact:hover {
            background-color: #001bb5;
        }

    /* Liên kết "Quên mật khẩu?" */
    .form_acc a {
        display: block;
        text-align: center;
        margin-top: 10px;
        font-size: 14px;
        color: #0026ff;
        text-decoration: none;
    }

        .form_acc a:hover {
            text-decoration: underline;
        }

    /* Chỉnh sửa CAPTCHA */
    .uk-text-center {
        text-align: center;
        margin-top: 10px;
    }
</style>
        <div class="sidebar_icon" style="display: none;">
            <ul class="list-123_icon" style="list-style: none;">
                <li style="padding-bottom: 20px;">
                    <a href="/DichVuTrucTuyen/DichVuCapDienMoi">
                        <div align="right">
                            <img id="imgCapDien" src="../Content/BCA_POPUP/icon_popup/Web1.png">
                        </div>
                    </a>
                </li>
                <li style="padding-bottom: 20px;">
                    <a href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=7">
                        <div align="right">
                            <img id="imgNgungCapDien" src="../Content/BCA_POPUP/icon_popup/Web2.png">
                        </div>
                    </a>
                </li>
                <li style="padding-bottom: 20px;">
                    <a href="/DichVuTTCSKH/IndexNPC?index=2">
                        <div align="right">
                            <img id="imgTraCuu" src="../Content/BCA_POPUP/icon_popup/Web3.png">
                        </div>
                    </a>
                </li>
                <li style="padding-bottom: 20px;">
                    <a href="/DichVuTTCSKH/DichVuTTCSKHNPC?index=8">
                        <div align="right">
                            <img id="imgnGhiChiSo" src="../Content/BCA_POPUP/icon_popup/Web4.png">
                        </div>
                    </a>
                </li>
                <li style="padding-bottom: 20px;">
                    <a href="/DichVuTrucTuyen/ThanhToanTrucTuyenSPC_TTTT">
                        <div align="right">
                            <img id="imgThanhToan" src="../Content/BCA_POPUP/icon_popup/Web5.png">
                        </div>
                    </a>
                </li>
            </ul>
        </div>

         
<footer id="footer-uk" class="uk-text-white">
    <div class="uk-section-small">
        <div class="uk-container">
            <div class="uk-child-width-1-3@m uk-flex-between" uk-grid>
                <div>
                    <address>
                        Trung tâm Chăm sóc khách hàng  <br>
                        Tổng Công ty Điện lực miền Bắc<br>
                        Số giấy phép: 0100100417-048 cấp ngày 21/10/2015
                    </address>
                    <a href="http://online.gov.vn/HomePage/CustomWebsiteDisplay.aspx?DocId=60326" target="_blank"><img src="/Content/BCA/imgs/BCT_NEW.png" alt=""></a>
                </div>
                <div>
                    <address>
                        <ul class="uk-list list1">
                            <li>
                                <span>Địa chỉ: </span>
                                Thửa số 2, Lô VP1, Khu bán đảo Linh Đàm, Phường Hoàng Liệt, Thành phố Hà Nội.
                            </li>
                            <li>
                                <span>Hotline:</span>
                                1900 6769
                            </li>
                            <li>
                                <span>Email:</span>
                                cskh@npc.com.vn
                            </li>
                        </ul>
                    </address>
                </div>
                <div>
                    <ul class="uk-list uk-column-1-2 list2">
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/Home/GioiThieuNPC">Giới thi&#234;̣u</a>
                        </li>
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/Home/HuongDanThuTucNPC">Hướng dẫn</a>
                        </li>
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/DichVuTrucTuyen/DichVuCapDienMoi">Dịch vụ trực tuyến</a>
                        </li>
                        
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/DichVuTTCSKH/IndexNPC">Tra cứu</a>
                        </li>
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/Home/TinMoiEVN">Tin tức</a>
                        </li>
                        <li class="uk-position-relative">
                            <i class="fa fa-angle-right uk-position-center-left" aria-hidden="true"></i>
                            <a href="/Home/LienHeNPC">Li&#234;n hệ</a>
                        </li>
                        
                        <li>
                            
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
    <div class="bottom">
        <div class="uk-container text-center">© 2025 EVNNPC CSKH. ALL RIGHTS RESERVED.</div>
    </div>
    <!-- Messenger Plugin chat Code -->
    <div id="fb-root"></div>

    <!-- Your Plugin chat code -->
    <div id="fb-customer-chat" class="fb-customerchat">
    </div>

    <script>
      var chatbox = document.getElementById('fb-customer-chat');
      chatbox.setAttribute("page_id", "543357612668147");
      chatbox.setAttribute("attribution", "biz_inbox");

      window.fbAsyncInit = function() {
        FB.init({
          xfbml            : true,
          version          : 'v12.0'
        });
      };

      (function(d, s, id) {
        var js, fjs = d.getElementsByTagName(s)[0];
        if (d.getElementById(id)) return;
        js = d.createElement(s); js.id = id;
        js.src = 'https://connect.facebook.net/vi_VN/sdk/xfbml.customerchat.js';
        fjs.parentNode.insertBefore(js, fjs);
      }(document, 'script', 'facebook-jssdk'));
    </script>
    
    <!-- Notification bar -->
    
</footer>


        <div class="back-to-top">
            <a href="javascript:void(0);">L&#234;n</a>
        </div>
    </section>
    <!-- Footer Scripts -->

    
    

    <script>
                    function send_data_DVC() {
                        var URL = "https://xacthuc.dichvucong.gov.vn/oauth2/authorize?response_type=code&client_id=jZOuLMbu8nhEADmkysKHY3kPUXIa&redirect_uri=https://cskh.npc.com.vn/home/XacthucDVC&scope=openid&acr_values=LoA1";
                        window.location.href = URL;
                    }
                    function frmAccount_OnclickDangKy() {
                        var tmpTenDN = KiemTraTenDangNhap($("#TenDangNhap_DK").val());
                        if (tmpTenDN == "OK") {
                            $('#frmAccount_DangKy').submit();
                        }
                        else {
                            var span = document.getElementById("frmAccount_ThongBaoLoi");
                            span.textContent = tmpTenDN;
                        }

                        //$('#frmAccount_DangKy').submit();
                    }


                    function KiemTraTenDangNhap(TenDangNhap) {
                        var strKetQua = "";
                        var str = TenDangNhap;
                        var regex = /^([a-zA-Z0-9@_.])+$/;
                        if (regex.test(str)) {
                            if (str.length >= 6) {
                                var pattern = /@{2,}/;
                                if (pattern.test(str)) {
                                    strKetQua = "Tên đăng nhập chỉ chứa 1 ký tự @.";
                                }
                                else {
                                    strKetQua = "OK";
                                }

                            }
                            else {
                                strKetQua = "Tên đăng nhập tối thiểu 6 ký tự.";

                            }

                            //isValid = true;
                        }
                        else {
                            strKetQua = "Tên đăng nhập sử dụng a-z 0-9.";
                        }
                        //alert(strKetQua);
                        return strKetQua;

                    }
                    function KiemTraMaKhachHang(MaKhachHang) {
                        var isValid = false;
                        var regex = /^([p,P][b,B,k,K][0-9]{11})+$/;
                        if (regex.test(MaKhachHang)) {
                            isValid = true;
                        }
                        return isValid;
                    }
                    function frmDangKy_OnkeyUp_NhapLaiMatKhau() {
                        var tmpValue = document.getElementById("NhapLaiMatKhau_DangKy");
                        tmpValue.value = tmpValue.value.trim();
                    }
                    function frmDangKy_OnkeyUp_MatKhau() {
                        var tmpValue = document.getElementById("MatKhau_DangKy");
                        tmpValue.value = tmpValue.value.trim();
                    }
                    function frmDangKy_OnkeyUp() {
                        var tmpValue = document.getElementById("MaKhachHang_DK");
                        tmpValue.value = tmpValue.value.toUpperCase().trim();
                    }
                    function frmDangKy_OnkeyUp_MKH() {
                        var tmpValue = document.getElementById("TenDangNhap_DK");
                        tmpValue.value = tmpValue.value.trim();
                    }
                    function frmDangNhap_OnkeyUp_MKH_DN() {
                        var tmpValue = document.getElementById("frmDangKy_TenDangNhap_DN");
                        tmpValue.value = tmpValue.value.trim();
                    }
                    function frmDangNhap_OnkeyUp_MKH() {
                        var tmpValue = document.getElementById("frmDangKy_MatKhau_DN");
                        tmpValue.value = tmpValue.value.trim();
                    }

                    $(".toggle-password").click(function () {
                        $(this).toggleClass("fa-eye fa-eye-slash");
                        var input = $($(this).attr("toggle"));
                        if (input.attr("type") == "password") {
                            input.attr("type", "text");
                        } else {
                            input.attr("type", "password");
                        }
                    });

    </script>

    <script>
        $(document).ready(function () {
            document.getElementById("imgCapDien").style.borderRadius = "30px";
            document.getElementById("imgCapDien").style.boxShadow = "5px 10px 10px #888888";
            document.getElementById("imgNgungCapDien").style.borderRadius = "30px";
            document.getElementById("imgNgungCapDien").style.boxShadow = "5px 10px 10px #888888";
            document.getElementById("imgTraCuu").style.borderRadius = "30px";
            document.getElementById("imgTraCuu").style.boxShadow = "5px 10px 10px #888888";
            document.getElementById("imgnGhiChiSo").style.borderRadius = "30px";
            document.getElementById("imgnGhiChiSo").style.boxShadow = "5px 10px 10px #888888";
            document.getElementById("imgThanhToan").style.borderRadius = "30px";
            document.getElementById("imgThanhToan").style.boxShadow = "5px 10px 10px #888888";
        });
        $(document).ready(function () {
            $("#imgCapDien").hover(
                function () {
                    document.getElementById("imgCapDien").src = "/Content/BCA_POPUP/icon_popup/Web6.png";
                    document.getElementById("imgCapDien").style.boxShadow = "5px 10px 10px #888888";
                    document.getElementById("imgCapDien").style.borderRadius = "30px";
                },
                function () {
                    document.getElementById("imgCapDien").src = "/Content/BCA_POPUP/icon_popup/Web1.png";
                    document.getElementById("imgCapDien").style.boxShadow = "5px 10px 10px #888888";
                }
            );
        });
        $(document).ready(function () {
            $("#imgNgungCapDien").hover(
                function () {
                    document.getElementById("imgNgungCapDien").src = "/Content/BCA_POPUP/icon_popup/Web7.png";
                    document.getElementById("imgNgungCapDien").style.boxShadow = "5px 10px 10px #888888";
                    document.getElementById("imgNgungCapDien").style.borderRadius = "30px";
                },
                function () {
                    document.getElementById("imgNgungCapDien").src = "/Content/BCA_POPUP/icon_popup/Web2.png";
                    document.getElementById("imgNgungCapDien").style.boxShadow = "5px 10px 10px #888888";
                }
            );
        });
        $(document).ready(function () {
            $("#imgTraCuu").hover(
                function () {
                    document.getElementById("imgTraCuu").src = "/Content/BCA_POPUP/icon_popup/Web8.png";
                    document.getElementById("imgTraCuu").style.boxShadow = "5px 10px 10px #888888";
                    document.getElementById("imgTraCuu").style.borderRadius = "30px";
                },
                function () {
                    document.getElementById("imgTraCuu").src = "/Content/BCA_POPUP/icon_popup/Web3.png";
                    document.getElementById("imgTraCuu").style.boxShadow = "5px 10px 10px #888888";
                }
            );
        });
        $(document).ready(function () {
            $("#imgnGhiChiSo").hover(
                function () {
                    document.getElementById("imgnGhiChiSo").src = "/Content/BCA_POPUP/icon_popup/Web9.png";
                    document.getElementById("imgnGhiChiSo").style.boxShadow = "5px 10px 10px #888888";
                    document.getElementById("imgnGhiChiSo").style.borderRadius = "30px";
                },
                function () {
                    document.getElementById("imgnGhiChiSo").src = "/Content/BCA_POPUP/icon_popup/Web4.png";
                    document.getElementById("imgnGhiChiSo").style.boxShadow = "5px 10px 10px #888888";
                }
            );
        });
        $(document).ready(function () {
            $("#imgThanhToan").hover(
                function () {
                    document.getElementById("imgThanhToan").src = "/Content/BCA_POPUP/icon_popup/Web10.png";
                    document.getElementById("imgThanhToan").style.boxShadow = "5px 10px 10px #888888";
                    document.getElementById("imgThanhToan").style.borderRadius = "30px";
                },
                function () {
                    document.getElementById("imgThanhToan").src = "/Content/BCA_POPUP/icon_popup/Web5.png";
                    document.getElementById("imgThanhToan").style.boxShadow = "5px 10px 10px #888888";
                }
            );
        });
    </script>

    
</body>
</html>
