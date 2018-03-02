function dothis() {
	alert("Taa Daa");
}

function confirmPaymentDone(result) {
	console.log("Result: ");
	console.log(result);

    if(result.status != 200 && result.errors.length > 0) {
        alert("do something because this error");
        return false;
    }

    alert("do something when this success");
}

$(document).ready(function () {
	$("#this-btn").click(function() {
		var request = $.ajax({
            url:    "/ajax/payment_page?type=success",
            method: "POST",
            dataType: "json",
        });
        request.done(confirmPaymentDone);
	});
});