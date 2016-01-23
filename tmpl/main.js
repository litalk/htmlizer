//<![CDATA[
function toggleDisplay(elementId) {
	var elm = document.getElementById(elementId + 'error');
	if (elm && typeof elm.style != "undefined") {
		if (elm.style.display == "none") {
			elm.style.display = "";
			document.getElementById(elementId + 'off').style.display = "none";
			document.getElementById(elementId + 'on').style.display = "inline";
		} else if (elm.style.display == "") {
			elm.style.display = "none";
			document.getElementById(elementId + 'off').style.display = "inline";
			document.getElementById(elementId + 'on').style.display = "none";
		}
	}
}
//]]>

function initPreWidth() {
	var tds =  document.getElementsByName('testcase_name');
	if (tds.length <= 0) {
		return;
	}
	
	var width = tds[0].offsetWidth + "px";
	var pres = document.getElementsByTagName('pre');
	for (var i = 0 ; i < pres.length ; ++i) {
		pres[i].style.width = width;
	}
}