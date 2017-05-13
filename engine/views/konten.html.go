// Code generated by hero.
// source: E:\appdanlib\Go_project\src\github.com\edot92\gotherimg\views\konten.html
// DO NOT EDIT!
package template

import "bytes"

func HalamanRender(dataInterface map[string]string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>

<!--
Website System 	  : Thermal Imaging Client Coluding
Start Develop on  : 13 Mei 2017
Develop by		  : Edy Prasetiyo
Contact Developer : edyprasetiyoo@gmail.com
Facebook          : https://www.facebook.com/eddot.fu
Phone 			  : 082210160003 & WA: 083876989317
-->

<html lang="en">

	<head>

		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta charset="utf-8" />
		<title> SAMPLE</title>
		<meta name="_xsrf" id="_xsrf" content="" />
		<meta name="description" content="top menu &amp; navigation" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />

		`)
	buffer.WriteString(`<!-- bootstrap & fontawesome -->
<link rel="stylesheet" href="static/css/bootstrap.min.css" />
<link rel="stylesheet" href="static/font-awesome/4.5.0/css/font-awesome.min.css" />

<!-- page specific plugin styles -->

<!-- text fonts -->
<link rel="stylesheet" href="static/css/fonts.googleapis.com.css" />

<!-- ace styles -->
<link rel="stylesheet" href="static/css/ace.min.css?v='" + Date.now()+ "'" class="ace-main-stylesheet" id="main-ace-style"
/>

<!--[if lte IE 9]>
			<link rel="stylesheet" href="static/css/ace-part2.min.css" class="ace-main-stylesheet" />
		<![endif]-->
<link rel="stylesheet" href="static/css/ace-skins.min.css" />
<link rel="stylesheet" href="static/css/ace-rtl.min.css" />

<!--[if lte IE 9]>
		  <link rel="stylesheet" href="static/css/ace-ie.min.css" />
		<![endif]-->

<!-- inline styles related to this page -->



<!-- HTML5shiv and Respond.js for IE8 to support HTML5 elements and media queries -->

<!--[if lte IE 8]>
		<script src="static/js/html5shiv.min.js"></script>
		<script src="static/js/respond.min.js"></script>
		<![endif]-->

<link rel="stylesheet" type="" href="static/css/select2.min.css">
<link rel="stylesheet" type="" href="static/css/jqx.base.css">
<link rel="stylesheet" type="" href="static/css/sweetalert2.min.css">
<link rel="stylesheet" type="" href="static/css/bootstrap-datetimepicker.min.css">
<link rel="stylesheet" type="" href="static/css/bootstrap-table.css">
<!--<link rel="stylesheet" type="" href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/css/select2.min.css">
<link rel="stylesheet" type="" href="http://www.jqwidgets.com/jquery-widgets-demo/jqwidgets/styles/jqx.base.css">
<link rel="stylesheet" type="" href="static/css/sweetalert2.min.css">
<link rel="stylesheet" type="" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datetimepicker/4.17.47/css/bootstrap-datetimepicker.min.css">

<link rel="stylesheet" type="" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.11.1/bootstrap-table.css">-->

<!--custom-->
<style>
	.header {
		margin-top: 0px !important;
	}
	/*.sidebar+.main-content {
		margin-left: 250px !important;
	}

	.sidebar {
		width: 250px !important;
	}*/

	@media (min-width: 768px) {
		/*.select2-container--bootstrap {
			width: 34px !important;
		}*/
	}

	@media (min-width: 992px) {}

	@media (min-width: 1200px) {}

	.modal-dialog {
		position: relative;
		width: auto;
		max-width: 600px;
		margin: 10px;
	}

	.modal-sm {
		max-width: 300px;
	}

	.modal-lg {
		max-width: 900px;
	}

	@media (min-width: 768px) {
		.modal-dialog {
			margin: 30px auto;
		}
	}

	@media (min-width: 320px) {
		.modal-sm {
			margin-right: auto;
			margin-left: auto;
		}
	}

	@media (min-width: 620px) {
		.modal-dialog {
			margin-right: auto;
			margin-left: auto;
		}
		.modal-lg {
			margin-right: 10px;
			margin-left: 10px;
		}
	}

	@media (min-width: 920px) {
		.modal-lg {
			margin-right: auto;
			margin-left: auto;
		}
	}
</style>`)

	buffer.WriteString(`<!-- basic scripts -->

<!--[if !IE]> -->
<!--script src="static/js/jquery-2.1.4.min.js">
    </script>-->

<!-- <![endif]-->
<!--<script src="static/js/jquery-1.11.3.min.js"></script>-->

<!--[if IE]>-->
<script src="static/js/jquery.min.js"></script>
<!--<![endif]-->
<script type="text/javascript">
    if ('ontouchstart' in document.documentElement) document.write("<script src='static/js/jquery.mobile.custom.min.js'>" + "<" + "/script>");

</script>
<script src="static/js/bootstrap.min.js"></script>

<!-- page specific plugin scripts -->

<!-- ace scripts -->
<script src="static/js/ace-elements.min.js"></script>
<script src="static/js/ace.min.js"></script>
<!--<script src="http://ace.jeka.by/assets/js/ace-extra.min.js"></script>-->
<!-- ace settings handler -->
<script src="static/js/ace-extra.min.js"></script>
<script src="static/js/select2.min.js"></script>
<script src="static/js/jqxcore.js"></script>
<script src="static/js/jqx-all.js"></script>
<script src="static/js/jquery.serializejson.min.js"></script>
<script src="static/js/axios.min.js"></script>
<script src="static/js/js.cookie.min.js"></script>
<script src="static/js/moment-with-locales.min.js"></script>
<script src="static/js/id.js"></script>
<script src="static/js/sweetalert2.min.js"></script>
<script src="static/js/bootstrap-datetimepicker.min.js"></script>
<script src="static/js/bootstrap-notify.min.js"></script>
<script src="static/js/jquery.inputmask.bundle.js"></script>
<script src="static/js/bootstrap-table.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/js/select2.min.js"></script>
<!--<script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.3/js/i18n/id.js"></script>-->

<!--<script src="http://www.jqwidgets.com/jquery-widgets-demo/jqwidgets/jqxcore.js"></script>
<script src="static/js/jqx-all.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.serializeJSON/2.8.1/jquery.serializejson.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/js-cookie/2.1.4/js.cookie.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/moment-with-locales.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/locale/id.js"></script>
<script src="https://cdn.jsdelivr.net/sweetalert2/6.6.2/sweetalert2.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datetimepicker/4.17.47/js/bootstrap-datetimepicker.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/mouse0270-bootstrap-notify/3.1.7/bootstrap-notify.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/jquery.inputmask.bundle.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-table/1.11.1/bootstrap-table.js"></script>-->
<!--<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.regex.extensions.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.extensions.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.numeric.extensions.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.date.extensions.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask.phone.extensions.js"></script>-->
<!--<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/phone-codes/phone.min.js"></script>-->
<!--<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/phone-codes/phone-be.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/inputmask/inputmask/phone-codes/phone-ru.js"></script>-->
<!--<script scr="https://cdnjs.cloudflare.com/ajax/libs/jquery.inputmask/3.3.4/jquery.inputmask.bundle.min.js"></script>-->
<!--<script src="https://rawgit.com/wenzhixin/bootstrap-table/master/src/bootstrap-table.js"></script>
<script src="https://cdn.datatables.net/1.10.15/css/jquery.dataTables.min.js"></script>-->


<script>
    // moment.locale('id');
    // moment.updateLocale('id', null)
    // moment.updateLocale('id', null);
    // $.fn.select2.defaults.set("theme", "bootstrap");

</script>`)

	buffer.WriteString(`
				<!-- Latest compiled and minified CSS -->


				<style>
					.header-site {
						background: url("static/header.jpg");
						/* Full height */
						/*height: 100%;*/
						/* Center and scale the image nicely */
						/*background-position: center;
					background-repeat: no-repeat;
					background-size: cover;*/
					}

					div#fixedheader {
						position: fixed;
						/*top: 0px;*/
						/*left: 0px;*/
						/*width: 100%;*/
						/*height: 100%;*/
						color: #CCC;
						background: #333;
						/*padding: 20px;*/
					}

					.navbar {
						/*background-color: darkorange !important;*/
					}
				</style>
				`)

	buffer.WriteString(`


	</head>

	<body class="no-skin">
		`)
	buffer.WriteString(`<div id="navbar" class="navbar navbar-default          ace-save-state">
    <div class="navbar-container ace-save-state" id="navbar-container">
        <button type="button" class="navbar-toggle menu-toggler pull-left" id="menu-toggler" data-target="#sidebar">
					<span class="sr-only">Toggle sidebar</span>

					<span class="icon-bar"></span>

					<span class="icon-bar"></span>

					<span class="icon-bar"></span>
				</button>

        <div class="navbar-header pull-left">
            <a href="#" class="navbar-brand">
						<small>
							<i class="fa fa-cloud"></i>
							IMAGE THERMAL CLIENT CLOUD
						</small>
					</a>
        </div>


    </div>
    <!-- /.navbar-container -->
</div>`)

	buffer.WriteString(`

			<konten>
				<div class="main-container ace-save-state" id="main-container">
					`)
	buffer.WriteString(`<div id="sidebar" class="sidebar                  responsive                    ace-save-state" data-sidebar="true" data-sidebar-scroll="true"
    data-sidebar-hover="true">
    <script type="text/javascript">
        try { ace.settings.loadState('sidebar') } catch (e) { }
    </script>

    <!-- /.sidebar-shortcuts -->

    <ul class="nav nav-list" id="navbarmenu1" style="top: 0px;">
        <li class="highlight ">
            <a href="#sumber_gambar" class="dropdown-toggle">
        <span class="menu-icon"><i class="menu-icon fa fa-pencil-square-o"></i></span>
        <span class="menu-text">Sumber Gambar</span>
    </a>
            <b class="arrow"></b>
        </li>
        <li class="highlight ">
            <a href="#logs_konten" class="dropdown-toggle">
        <span class="menu-icon"><i class="menu-icon fa fa-bell-o"></i></span>
        <span class="menu-text">Logs Cloud Dan Alarm</span>
    </a>
            <b class="arrow"></b>
        </li>
        <li class="highlight ">
            <a href="#pengaturan_cloud" class="dropdown-toggle">
        <span class="menu-icon"><i class="menu-icon fa fa-wrench"></i></span>
        <span class="menu-text">Pengaturan CLoud</span>
    </a>
            <b class="arrow"></b>
        </li>
        <li class="highlight hover hsub">
            <a href="#logout" class="dropdown-toggle">
        <span class="menu-icon"><i class="fa fa-lock"></i></span>
        <span class="menu-text">Logout</span>
    </a>
            <b class="arrow"></b>
        </li>

    </ul>
    <!-- /.nav-list -->

    <div class="sidebar-toggle sidebar-collapse" id="sidebar-collapse">
        <i id="sidebar-toggle-icon" class="ace-icon fa fa-angle-double-left ace-save-state" data-icon1="ace-icon fa fa-angle-double-left"
            data-icon2="ace-icon fa fa-angle-double-right"></i>
    </div>
</div>`)

	buffer.WriteString(`
                <div class="main-content">
                    <div class="main-content-inner">
                        <div class="page-content">
                            `)
	buffer.WriteString(`<!--start logs konten-->
<div id="sumber_gambar">
    <script>
        function point_it(event) {
            pos_x = (event.offsetX ? (event.offsetX) : event.pageX - document.getElementById("pointer_div").offsetLeft);
            pos_y = (event.offsetY ? (event.offsetY) : event.pageY - document.getElementById("pointer_div").offsetTop);
            $('#cross').css('left', (pos_x - 1));
            $('#cross').css('top', (pos_y - 15));
            $('#cross').css('visibility', "visible");
            paramx = (pos_x / 3).toFixed(0);
            paramy = (pos_y / 3).toFixed(0);
            $('#form_x').attr('value', paramx + "pixel");
            $('#form_y').attr('value', paramy + "pixel");
            requestColor(paramx, paramy)
        }

        function requestColor(paramx, paramy) {

            paramUrl = "gambar/" + paramx + "/" + paramy;
            axios(paramUrl, {
                responseType: 'json', // default
            }).then(function (response) {
                console.log(response);

                $('#hex').html("HEX COLOR:" + response.data.Hex)
                $('#rgb').html("RGB:" + JSON.stringify(response.data.R) + " ," + JSON.stringify(response.data.G) + " ," + JSON.stringify(response.data.B))
                $('#suhu').html("Suhu:" + JSON.stringify(response.data.Suhu))

            })
                .catch(function (error) {
                    console.log(error);
                    alert(error);
                });
        }
    </script>
    <div class="row">
        <h4 class="header smaller lighter blue">
            Pengaturan Sumber Gambar
        </h4>
    </div>
    <div class="row">
        <div class="col-md-3 col-sm-12">
            <div class="panel panel-primary" style="width: 242px;">
                <div class="panel-heading">Sumber Gambar
                    <a href="#" data-action="reload" style="color: white;"><i class="ace-icon fa fa-refresh"></i></a>
                </div>
                <div class="panel-body" style="padding-top: 15px;padding-left: 0px;">
                    <div class="form-horizontal">
                        <form name="pointform" method="post">
                            <div id="pointer_div" onclick="point_it(event)" width='100%' height='100%' style="  background-size:cover;background-repeat:no-repeat;background-position: center center;background-image:url('static/tes.jpg');width:240px;height:240px;">
                                <img src="static/point.gif" id="cross" style="position:relative;visibility:hidden;z-index:2;"></img>
                        </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-md-3 col-sm-12">
            <div class="panel panel-primary">
                <div class="panel-heading">Nilai Real Point Gambar</div>
                <div class="panel-body">
                    Pointer x = <input type="text" id="form_x" name="form_x" size="8" />
                    <br> Pointer y = <input type="text" id="form_y" name="form_y" size="8" />
                    <br>
                    <pre id="hex"></pre>
                    <pre id="rgb"></pre>
                    <pre id="suhu"></pre>
                    <br>
                    <button type="" class="btn btn-primary"> Tambahkan ke alarm tabel </button>
                </div>
            </div>
        </div>
        <div class="col-md-6 col-sm-12">
            <div class="panel panel-primary">
                <div class="panel-heading">Tabel Setting Alarm</div>
                <div class="panel-body">
                    <div id="tabelSetTemp">
                        <!--isi table-->
                    </div>
                </div>
            </div>
        </div>
    </div>


    <script>
        $(document).ready(function () {
            var dataTes = [{ "id": 1, "pointer_x": "20", "pointer_y": "30", "value_min_alarm": "10", "value_max_alarm": "10" }];
            inittableJqx(dataTes);
            function inittableJqx(param) {
                // prepare the data
                var source =
                    {
                        localdata: param,
                        datatype: "json",
                        datafields: [
                            { name: 'id', type: 'int' },
                            { name: 'pointer_x', type: 'string' },
                            { name: 'pointer_y', type: 'int' },
                            // { name: 'set_temperature', type: 'string' },
                            { name: 'value_min_alarm', type: 'string' },
                            { name: 'value_max_alarm', type: 'string' },
                        ],
                        id: 'id',
                        // url: url
                    };
                var dataAdapter = new $.jqx.dataAdapter(source);
                $("#tabelSetTemp").jqxGrid(
                    {
                        width: "100%",
                        source: dataAdapter,
                        columnsresize: true,
                        columns: [
                            { text: 'id', datafield: 'id', width: 10, cellsalign: "center" },
                            { text: 'Pointer X', datafield: 'pointer_x', minwidth: 70, cellsalign: "center" },
                            { text: 'Pointer Y', datafield: 'pointer_y', minwidth: 70, cellsalign: "center" },
                            // { text: 'Set Temperature', datafield: 'set_temperature', width: 50 },
                            { text: 'Min Alarm &#8451;', datafield: 'value_min_alarm', minwidth: 50, cellsalign: "center" },
                            { text: 'Max Alarm &#8451;', datafield: 'value_max_alarm', minwidth: 50, cellsalign: "center" },
                            {
                                text: 'Edit', minwidth: 120, datafield: 'Edit', columntype: 'button', cellsrenderer: function () {
                                    return "Edit";
                                }, buttonclick: function (row) {
                                    editrow = row;
                                    var offset = $("#tabelSetTemp").offset();
                                    var dataRecord = $("#tabelSetTemp").jqxGrid('getrowdata', editrow);
                                }
                            },
                            {
                                text: 'Hapus', minwidth: 120, datafield: 'Hapus', columntype: 'button', cellsrenderer: function () {
                                    return "Hapus";
                                }, buttonclick: function (row) {
                                    editrow = row;
                                    var offset = $("#tabelSetTemp").offset();
                                    var dataRecord = $("#tabelSetTemp").jqxGrid('getrowdata', editrow);
                                }
                            }
                        ]
                    });

            }
        })
    </script>
</div>`)

	buffer.WriteString(`<!--start logs konten-->
<div id="logs_konten">

    <div class="row">
        <h4 class="header smaller lighter blue">
            Logs Device Cloud Dan Alarm
        </h4>
    </div>
    <div class="row">
        <div class="col-md-4">
            <div class="form-group">
                <label for="" class="control-label col-sm-2 col-md-2" style="padding-top: 5px;">Status</label>
                <div class="col-sm-9">
                    <select class="form-control">
    <option value="" selected>Pengiriman Alarm</option>
    <option value="">Pengiriman Berhasil</option>
    <option value="">Pengiriman Gagal</option>
</select>
                </div>
            </div>
        </div>
        <div class="col-md-3">
            <div class="form-group">
                <label class="control-label col-sm-2 col-md-2" for="" style="padding-top: 5px;">Dari </label>
                <div class="col-sm-9">
                    <input type="text" id="InputTahananBaru_TanggalMasuk" name="TanggalMasuk" placeholder="DD/MM/YYYY" class="datepicker form-control">
                </div>
            </div>
        </div>
        <div class="col-md-3">
            <div class="form-group">
                <label class="control-label col-sm-2 col-md-2" for="" style="padding-top: 5px;">s/d</label>
                <div class="col-sm-9">
                    <input type="text" id="InputTahananBaru_TanggalMasuk" name="TanggalMasuk" placeholder="DD/MM/YYYY" class="datepicker form-control">
                </div>
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-md-12">
            <div class="tabelLogs">
            </div>
        </div>
    </div>

    <script>
        $(document).ready(function () {
            var dataTes = [{ "id": 1, "pointer_x": "20", "pointer_y": "30", "value_min_alarm": "10", "value_max_alarm": "10" }];
            inittableJqx(dataTes);
            function inittableJqx(param) {
                // prepare the data
                var source =
                    {
                        localdata: param,
                        datatype: "json",
                        datafields: [
                            { name: 'id', type: 'int' },
                            { name: 'pointer_x', type: 'string' },
                            { name: 'pointer_y', type: 'int' },
                            // { name: 'set_temperature', type: 'string' },
                            { name: 'value_min_alarm', type: 'string' },
                            { name: 'value_max_alarm', type: 'string' },
                        ],
                        id: 'id',
                        // url: url
                    };
                var dataAdapter = new $.jqx.dataAdapter(source);
                $("#tabelSetTemp").jqxGrid(
                    {
                        width: "100%",
                        source: dataAdapter,
                        columnsresize: true,
                        columns: [
                            { text: 'id', datafield: 'id', width: 10, cellsalign: "center" },
                            { text: 'Pointer X', datafield: 'pointer_x', minwidth: 70, cellsalign: "center" },
                            { text: 'Pointer Y', datafield: 'pointer_y', minwidth: 70, cellsalign: "center" },
                            // { text: 'Set Temperature', datafield: 'set_temperature', width: 50 },
                            { text: 'Min Alarm &#8451;', datafield: 'value_min_alarm', minwidth: 50, cellsalign: "center" },
                            { text: 'Max Alarm &#8451;', datafield: 'value_max_alarm', minwidth: 50, cellsalign: "center" },
                            {
                                text: 'Edit', minwidth: 120, datafield: 'Edit', columntype: 'button', cellsrenderer: function () {
                                    return "Edit";
                                }, buttonclick: function (row) {
                                    editrow = row;
                                    var offset = $("#tabelSetTemp").offset();
                                    var dataRecord = $("#tabelSetTemp").jqxGrid('getrowdata', editrow);
                                }
                            },
                            {
                                text: 'Hapus', minwidth: 120, datafield: 'Hapus', columntype: 'button', cellsrenderer: function () {
                                    return "Hapus";
                                }, buttonclick: function (row) {
                                    editrow = row;
                                    var offset = $("#tabelSetTemp").offset();
                                    var dataRecord = $("#tabelSetTemp").jqxGrid('getrowdata', editrow);
                                }
                            }
                        ]
                    });

            }
        })
    </script>

</div>
<!--end logs konten-->`)

	buffer.WriteString(`<div id="pengaturan_cloud">
    <div class="row">
        <h4 class="header smaller lighter blue">
            Pengaturan Perangkat
        </h4>
    </div>
    <div class="row">
        <div class="col-md-6 col-sm-12">
            <div class="panel panel-primary">
                <div class="panel panel-heading">
                    Pengaturan Cloud
                </div>
                <div class="panel panel-body" style="padding: 15px;">
                    <form class="form-horizontal" action="#">
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="email">Server URL:</label>
                            <div class="col-sm-6">
                                <input type="text" class="form-control" id="email" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="pwd">Username:</label>
                            <div class="col-sm-6">
                                <input type="text" class="form-control" id="pwd" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="pwd">Password:</label>
                            <div class="col-sm-6">
                                <input type="password" class="form-control" id="pwd" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="pwd">ID Cloud:</label>
                            <div class="col-sm-6">
                                <input type="text" class="form-control" id="pwd" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="pwd">Enable Cloud:</label>
                            <div class="col-sm-6" style="margin-top: 5px;">
                                <label><input name="switch-field-1" class="ace ace-switch ace-switch-4" type="checkbox">
														<span class="lbl"></span></label>
                            </div>
                        </div>

                        <div class="form-group">
                            <div class="col-sm-offset-3 col-sm-12">
                                <button type="submit" class="btn btn-info"><i class="fa fa-link"></i>Test Koneksi</button>
                                <button type="submit" class="btn btn-primary"><i class="fa fa-save"></i>Simpan</button>

                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
        <div class="col-md-6 col-sm-12">
            <div class="panel panel-primary">
                <div class="panel panel-heading">
                    Pengaturan Modem
                </div>
                <div class="panel panel-body" style="padding: 15px;">
                    <form class="form-horizontal" action="#">
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Modem Selected :</label>
                            <div class="col-sm-8">
                                <div class="input-group input-group-sm">
                                    <select class="form-control">
                            <option value="" selected=""> HUAWEI ZTE</option>
                            </select>

                                    <span class="input-group-btn"><button type="button" class="btn btn-info">
																			<span class="ace-icon fa fa-refresh icon-on-right bigger-110"></span>                                    Refresh</button>
                                    </span><span class="input-group-btn">
    <button type="button" class="btn btn-primary">
																			<span class="ace-icon fa fa-save icon-on-right bigger-110">
                                                                                </span>SIMPAN</button>
                                    </span>
                                </div>

                            </div>

                        </div>
                        <div class="hr hr-24"></div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Username :</label>
                            <div class="col-sm-8">
                                <input type="text" class="form-control" id="" value="wap" placeholder="Enter">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Password :</label>
                            <div class="col-sm-8">
                                <input type="text" class="form-control" id="" value="wap123" placeholder="Enter">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">APN :</label>
                            <div class="col-sm-8">
                                <input type="text" class="form-control" id="" value="telkomesel" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Dial Number :</label>
                            <div class="col-sm-8">
                                <input type="text" class="form-control" id="" value="*99#" placeholder="Enter ">
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-offset-3 col-sm-12">
                                <button type="submit" class="btn btn-primary"><i class="fa fa-save"></i>Simpan Pengaturan APN</button>
                            </div>
                        </div>
                        <div class="hr hr-24"></div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Command Modem :</label>
                            <div class="col-sm-8">
                                <div class="input-group input-group-sm">
                                    <select class="form-control">
                            <option value="" selected="">Cek Pulsa</option>
                            <option value="" selected="">Cek Signal</option>
                            <option value="" selected="">Cek Internet</option>
                            </select>
                                    <span class="input-group-btn">
    <button type="button" class="btn btn-primary">
																			<span class="ace-icon fa fa-send icon-on-right bigger-110">
                                                                                </span>Kirim</button>
                                    </span>
                                </div>
                            </div>

                        </div>
                        <div class="form-group">
                            <label class="control-label col-sm-4" for="">Response :</label>
                            <div class="col-sm-8">
                                <textarea rows="4" cols="10" class="form-control"></textarea>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>`)

	buffer.WriteString(`
                                        <!-- /.page-content -->
                        </div>
                    </div>
                </div>
                `)
	buffer.WriteString(`<script type="text/javascript">
    var namaHalaman = ['sumber_gambar', 'logs_konten', 'pengaturan_cloud', 'settingreminder'];

    $(document).ready(function () {
        var token = 'pBBA8t6Rvve1Fv5fpd5Kv2TOX1T7wf68';
        $('#xrsf').attr('content', token);
        $('#xrsfInput').attr('value', token);
        for (var i = 0; i < namaHalaman.length; i++) {
            $('#' + namaHalaman[i]).css('display', 'none');
        }
        $("#navbarmenu1 li a").click(function (event) {

            Cookies.remove('tabaktif');
            var href_klikname = $(this).attr('href');
            event.preventDefault();
            if (href_klikname == "#logout") {
                document.cookie = 'jwt' + '=; expires=Thu, 01-Jan-70 00:00:01 GMT;';
                location.reload();
                return;
            }
            $("#navbarmenu1 .active").removeClass('active');
            $(this).parent().addClass('active');
            href_klikname = href_klikname.replace('#', '');
            console.log(href_klikname);
            status = gotoTab(href_klikname)
            if (status == true) {

                console.log("tes")
            }
        });
        var data = Cookies.get('tabaktif');
        var dataResFlashCokie = Cookies.getJSON('response_flash_json');
        // console.log("tabaktif =" + data);
        if (data == "inputtahananbaru") {
            gotoTab(data);

        }
        // console.log(dataResFlashCokie)
        if (dataResFlashCokie != undefined) {
            if (dataResFlashCokie.status == "success") {
                swal(
                    "Berhasil",
                    dataResFlashCokie.message,
                    "success"
                )
            } else if (dataResFlashCokie.status == "failed") {
                swal(
                    "Gagal",
                    dataResFlashCokie.message,
                    "error"
                )
            }

        }
        Cookies.remove('tabaktif');
        Cookies.remove('response_flash_json');

        function gotoTab(href_klikname) {
            var status = false;
            href_klikname = href_klikname.toLowerCase();
            for (var i = 0; i < namaHalaman.length; i++) {
                var namaHal = namaHalaman[i];
                if (href_klikname == namaHal) {
                    $('#' + namaHalaman[i]).css('display', 'inline');
                    status = true;
                } else {
                    $('#' + namaHalaman[i]).css('display', 'none');
                    // $('#navbarmenu1 > li').removeClass("active");

                }
            }
            return status;
        }
    });
                // jQuery(function ($) {

                // })

</script>


<script>
    $(document).ready(function () {
        $('.datepicker').datetimepicker({
            format: "DD-MM-YYYY",
            locale: 'id'
        });
        $(".datepicker").inputmask("99-99-9999", { "placeholder": "DD-MM-YYYY" });
    });

</script>`)

	buffer.WriteString(`
				</div>
			</konten>
			`)
	buffer.WriteString(`<!-- inline scripts related to this page -->
<script type="text/javascript">
    jQuery(function ($) {
        var $sidebar = $('.sidebar').eq(0);
        if (!$sidebar.hasClass('h-sidebar')) return;

        $(document).on('settings.ace.top_menu', function (ev, event_name, fixed) {
            if (event_name !== 'sidebar_fixed') return;

            var sidebar = $sidebar.get(0);
            var $window = $(window);

            //return if sidebar is not fixed or in mobile view mode
            var sidebar_vars = $sidebar.ace_sidebar('vars');
            if (!fixed || (sidebar_vars['mobile_view'] || sidebar_vars['collapsible'])) {
                $sidebar.removeClass('lower-highlight');
                //restore original, default marginTop
                sidebar.style.marginTop = '';

                $window.off('scroll.ace.top_menu')
                return;
            }


            var done = false;
            $window.on('scroll.ace.top_menu', function (e) {

                var scroll = $window.scrollTop();
                scroll = parseInt(scroll / 4);//move the menu up 1px for every 4px of document scrolling
                if (scroll > 17) scroll = 17;


                if (scroll > 16) {
                    if (!done) {
                        $sidebar.addClass('lower-highlight');
                        done = true;
                    }
                }
                else {
                    if (done) {
                        $sidebar.removeClass('lower-highlight');
                        done = false;
                    }
                }

                sidebar.style['marginTop'] = (17 - scroll) + 'px';
            }).triggerHandler('scroll.ace.top_menu');

        }).triggerHandler('settings.ace.top_menu', ['sidebar_fixed', $sidebar.hasClass('sidebar-fixed')]);

        $(window).on('resize.ace.top_menu', function () {
            $(document).triggerHandler('settings.ace.top_menu', ['sidebar_fixed', $sidebar.hasClass('sidebar-fixed')]);
        });


    });
    // https://junaidqadir.com/backyard/bootstrap_add_rmove_tabs/
    var pageImages = [];
    var pageNum = 1;
    /**
    * Reset numbering on tab buttons
    */
    function reNumberPages() {
        pageNum = 0;
        var tabCount = $('#pageTab > li').length;
        $('#pageTab > li').each(function () {
            var pageId = $(this).children('a').attr('href');
            if (pageId == "#page1") {
                return true;
            }
            pageNum++;
            $(this).children('a').html('Page ' + pageNum +
                '<button class="close" type="button" ' +
                'title="Remove this page">×</button>');
        });
    }

</script>`)

	buffer.WriteString(`
				<script>
					$('#sidebar-collapse').click(function (e) {
						alert("refresh");
						$('#tabelSetTemp').jqxGrid('refresh');
					})
				</script>
	</body>

</html>`)

}
