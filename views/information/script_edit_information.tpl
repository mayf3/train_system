{{define "information_script"}}
<script type="text/javascript">
$(document).ready(function(){
	$("button[name='checked_button0']").each(function(){
			$(this).html($("#member0").children('option:selected').html().substr(5));
			});
	$("button[name='checked_button1']").each(function(){
			$(this).html($("#member1").children('option:selected').html().substr(5));
			});
	$("button[name='checked_button2']").each(function(){
				$(this).html($("#member2").children('option:selected').html().substr(5));
			});

	$("#member0").change(function(){
		$("button[name='checked_button0']").each(function(){
				$(this).html($("#member0").children('option:selected').html().substr(5));
			});
		});
	$("#member1").change(function(){
		$("button[name='checked_button1']").each(function(){
				$(this).html($("#member1").children('option:selected').html().substr(5));
			});
		});
	$("#member2").change(function(){
		$("button[name='checked_button2']").each(function(){
				$(this).html($("#member2").children('option:selected').html().substr(5));
			});
		});

	$("button[name='checked_button0']").each(function(){
		if ($(this).prev().is(':checked')){
			$(this).prop("class", "btn btn-primary")
		}else{
			$(this).prop("class", "btn btn-default")
		}
		$(this).click(function(){
			if ($(this).prev().is(':checked')){
			 	$(this).prev().prop('checked', false);
				$(this).prop("class", "btn btn-default")
			} else {
			 	$(this).prev().prop('checked', true);
				$(this).prop("class", "btn btn-primary")
			}
			});
		});
	$("button[name='checked_button1']").each(function(){
		if ($(this).prev().is(':checked')){
			$(this).prop("class", "btn btn-primary")
		}else{
			$(this).prop("class", "btn btn-default")
		}
		$(this).click(function(){
			if ($(this).prev().is(':checked')){
			 	$(this).prev().prop('checked', false);
				$(this).prop("class", "btn btn-default")
			} else {
			 	$(this).prev().prop('checked', true);
				$(this).prop("class", "btn btn-primary")
			}
			});
		});
	$("button[name='checked_button2']").each(function(){
		if ($(this).prev().is(':checked')){
			$(this).prop("class", "btn btn-primary")
		}else{
			$(this).prop("class", "btn btn-default")
		}
		$(this).click(function(){
			if ($(this).prev().is(':checked')){
			 	$(this).prev().prop('checked', false);
				$(this).prop("class", "btn btn-default")
			} else {
			 	$(this).prev().prop('checked', true);
				$(this).prop("class", "btn btn-primary")
			}
			});
		});

	
	$("button[name='radio_button']").each(function(){
		if ($(this).prev().is(':checked')){
			$(this).prop("class", "btn btn-info")
		}else{
			$(this).prop("class", "btn btn-default")
		}
		console.log("read")
		$(this).click(function(){
			console.log("ok click")
			var tmp_name = $(this).prev().attr("name")
			console.log(tmp_name)
			var all = document.getElementsByName(tmp_name)
			if ($(this).prev().is(':checked') == false){
			 	$(this).prev().prop('checked', true);
			}
			console.log(all)
			for(i = 0; i < all.length; i++){
				if(all[i].checked == true){
					all[i].nextElementSibling.setAttribute("class", "btn btn-info")
				} else{
					all[i].nextElementSibling.setAttribute("class", "btn btn-default")
				}
			}
			});
		});
});
</script>
{{end}}}
