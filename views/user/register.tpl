{{template "base/base.tpl" .}}

{{define "meta"}}
	<title>Register - SYSU Train System</title>
{{end}}

{{define "body" }}
        <div class="container">
            <form class="form-horizontal col-sm-6 col-sm-offset-2" method="post" action="#">
                <div id="legend" class="">
                    <legend class="">Register</legend>
                </div>

                <div class="form-group">
                <!-- Search input-->
                    <label class="control-label col-sm-4"> Email : </label>
                    <div class="col-sm-8">
                        <input name="user[email]" placeholder="abc@gmail.com" class="form-control" type="text">
                    </div>
                </div>

                <div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4" for="input01"> Username : </label>
                    <div class="col-sm-8">
                        <input name="user[username]" placeholder="David" class="form-control" type="text">
                    </div>
                </div>

                <div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4" for="input01"> Password : </label>
                    <div class="col-sm-8">
                        <input name="user[password]" placeholder="placeholder" class="form-control" type="password">
                    </div>
                </div>
                
                <div class="form-group">
                <!-- Text input-->
                    <label class="control-label col-sm-4" for="input01"> Password Again: </label>
                    <div class="col-sm-8">
                        <input name="user[password_again]" placeholder="placeholder" class="form-control" type="password">
                    </div>
                </div>

                <div class="form-group">
                    <!-- Button -->
                    <label class="control-label col-sm-4"></label>
                    <div class="col-sm-8">
                        <button type="submit" class="btn btn-success">Sign up</button>
                    </div>
                </div>
            </form>
        </div>
{{end}}
