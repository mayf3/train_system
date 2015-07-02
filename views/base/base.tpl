<!DOCTYPE html>
<html>
    <head>
		{{template "base/head.tpl" .}}
    </head>
    <body role="document">
		<div class="container-fluid" id="wrapper">
			{{template "base/navbar.tpl" .}}
			<div id="main" class="container-fluid">
			{{template "body" .}}
			</div>
			<div class="wrapper-push"></div>
		</div>
		{{template "base/footer.tpl" .}}
	</body>
</html>
