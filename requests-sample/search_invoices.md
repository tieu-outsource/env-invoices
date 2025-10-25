## Request
await fetch("https://cskh.npc.com.vn/HoaDon/TraCuuHDSPC?ky=1&thang=9&nam=2025&_=1761322926877", {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0",
        "Accept": "text/html, */*; q=0.01",
        "Accept-Language": "en-US,en;q=0.5",
        "Content-Type": "application/json; charset=utf-8",
        "X-Requested-With": "XMLHttpRequest",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "same-origin",
        "Priority": "u=0"
    },
    "referrer": "https://cskh.npc.com.vn/DichVuTTCSKH/IndexNPC?index=2",
    "method": "GET",
    "mode": "cors"
});

## Response
<table class="uk-table uk-table-small uk-table-middle uk-table-divider table2">
    <thead>
        <tr>
            <th>STT</th>
            <th>M&#227; kh&#225;ch h&#224;ng</th>
            <th>ID h&#243;a đơn</th>
            <th>Seri-K&#253; hiệu</th>
            <th>Tổng Tiền</th>
            <th>Loại h&#243;a đơn</th>
            <th>Xem thể hiện</th>
        </tr>
    </thead>
    <tbody>
                    <tr>
                        <td>1</td>
                        <td class="uk-text-truncate">PA04GT7017040</td>
                        <td>1671732781</td>
                        <td class="uk-text-truncate">1046124</td>
                        <td>
28.386.839
                        </td>
                            <td>Tiền Điện</td>
                        <td>
                                <button 
                                        id="TraCuuHD_XemHoaDon"
                                        onclick="ChonThaoTac('XIq6RC1EGQud6DC___PGjwiQ==-PA04GT7017040-1-9-2025');"
                                        class="uk-button uk-button-primary btn-send-contact"
                                        style="min-width: inherit; padding: 0 10px; line-height: 32px;">
                                    <span uk-icon="search"></span>
                                </button>
                                <button class="uk-button uk-button-primary btn-send-contact"
                                        style="min-width: inherit; padding: 0 10px; line-height: 32px;"
                                        onclick="TraCuuHD_onClickTaiHoaDonNPC('XIq6RC1EGQud6DC___PGjwiQ==-PA04GT7017040');">
                                    <span uk-icon="download"></span>
                                </button>

                        </td>

                    </tr>

    </tbody>
</table>


<div id="frmTraCuu_ThongBaoHoaDon" class="uk-flex-top modal1_thanhtoan uk-modal" uk-modal="">
    <div class="uk-modal-dialog uk-modal-body uk-margin-auto-vertical uk-width-auto">
        <button class="uk-modal-close-default uk-close uk-icon" type="button" uk-close=""><svg width="14" height="14" viewBox="0 0 14 14" xmlns="http://www.w3.org/2000/svg" data-svg="close-icon"><line fill="none" stroke="#000" stroke-width="1.1" x1="1" y1="1" x2="13" y2="13"></line><line fill="none" stroke="#000" stroke-width="1.1" x1="13" y1="1" x2="1" y2="13"></line></svg></button>
        <div class="uk-margin" id="frmTraCuu_hdn_print">
            <button onclick="javascript: window.frames['frameIn'].focus(); parent['frameIn'].print();" class="uk-button uk-button-primary btn-send-contact"><span uk-icon="print" class="uk-icon"><svg width="20" height="20" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" data-svg="print"><polyline fill="none" stroke="#000" points="4.5 13.5 1.5 13.5 1.5 6.5 18.5 6.5 18.5 13.5 15.5 13.5"></polyline><polyline fill="none" stroke="#000" points="15.5 6.5 15.5 2.5 4.5 2.5 4.5 6.5"></polyline><rect fill="none" stroke="#000" width="11" height="6" x="4.5" y="11.5"></rect><rect width="8" height="1" x="6" y="13"></rect><rect width="8" height="1" x="6" y="15"></rect></svg></span> in thông tin hóa đơn</button>
        </div>
        <div class="uk-overflow-auto">
            <div id="frmHoaDon_HoaDonContent"  >
              <object id="frmTraCuuHoaDon_XemHD" style="width:100% ;" height="800" ></object>
            </div>
        </div>
    </div>
</div>




<form id="TraCuuHD_TempTaiHD">
    <input type="hidden" id="TraCuuHD_MA_KH" name="TraCuuHD_MA_KH" />
    <input type="hidden" id="TraCuuHD_IDHoaDon_MauMoi" name="TraCuuHD_IDHoaDon_MauMoi" />
    <input type="hidden" id="TraCuuHD_TempTaiHD_IDHoaDon" name="strid_hdon" />
    <input type="hidden" id="TraCuuHD_TempTaiHD_nam" name="nam" />
    <input type="hidden" id="TraCuuHD_TempTaiHD_thang" name="thang" />
    <input type="hidden" id="TraCuuHD_TempTaiHD_ky" name="ky" />
    <input type="hidden" id="TraCuuHD_TempTaiHD_DuongDan" name="TenThuMuc" />
        <input type="hidden" id="TraCuuHD_SoLuongHD" value="Tổng số h&#243;a đơn l&#224;  :<strong> 1  h&#243;a đơn</strong> " />
</form>

<div id="frmTraCuuHoaDon" class="uk-flex-top" uk-modal>
    <div class="uk-modal-dialog uk-modal-body uk-margin-auto-vertical">
        <div class="uk-margin box_a1">
            <h5 class="title1">
                Chọn thao tác
            </h5>
        </div>
        <div class="uk-overflow-auto">
            <div class="row" style="margin-left:0px;margin-right:0px;">
                <div class="col-md-12 col-lg-12 col-xs-12 col-sm-12 uk-margin" align="center">
                    <input class="uk-button uk-button-primary btn-send-contact btn-tracuu uk-icon"
                           onclick="XemHoaDonNPC();"
                           type="button" value="Xem hóa đon" />
                    <input class="uk-button uk-button-primary btn-send-contact btn-tracuu uk-icon"
                           onclick="XemChiTietNPC()"
                           type="button" value="Xem chi tiêt" />
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    function TraCuuHD_onClickTaiTatCaHoaDonSPC() {
        $('#TraCuuHD_TempTaiHD').attr("action", "/HoaDon/DownloadAllHD");
        $('#TraCuuHD_TempTaiHD').attr("method", "POST");
        $('#TraCuuHD_TempTaiHD').submit();
    }
</script>
