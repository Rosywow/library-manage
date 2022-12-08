<script lang="ts">
    import Button, { Label, Icon } from "@smui/button";
    import Textfield from "@smui/textfield";
    import Checkbox from "@smui/checkbox";
    import FormField from "@smui/form-field";
    import { push } from "svelte-spa-router";

    let username = "";
    let password = "";
    let checked = false;
    function Login() {
        let url = `api/login`;
        fetch(url, {
            method: "POST",
            mode: "no-cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: new Headers({
                "Content-Type": "application/json",
            }),
            redirect: "follow",
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        })
            .then((v) => {
                console.log(v);
                return v.json();
            })
            .then((res) => {
                if (res.status === 200) {
                    console.log("登录成功");
                    if (username === "root") {

                    } else {

                    }
                } else {
                    console.log("登录失败err:",res.err)
                }
            })
            .catch((err) => {
                console.log("err:", err);
            });
    }
</script>

<div class="container">
    <div class="image">
        <img src="\src\assets\icons8-signin-50 (1).png" alt="sign in" />
    </div>

    <h1>Sign in</h1>
    <div class="username">
        <Textfield
            class="shaped-outlined"
            variant="outlined"
            type="text"
            bind:value={username}
            label="Username"
        />
    </div>
    <div class="password">
        <Textfield
            class="shaped-outlined"
            variant="outlined"
            type="password"
            bind:value={password}
            label="Password"
        />
    </div>

    <div class="rememberme">
        <FormField>
            <Checkbox bind:checked />
            <span slot="label">Remember me.</span>
        </FormField>
    </div>

    <div class="button">
        <Button
            on:click={Login}
            style="background-color:#1976d2;color:#fff"
        >
            <Label>Sign in</Label>
        </Button>
    </div>

    <div class="others">
        <a href="#/">Forgot password?</a>
        <a href="#/signup">Don't have an account?Sign Up</a>
    </div>

    <footer>
        <p id="copyright" class="copyright text">
            Copyright &copy; 2022 by Rosy.
        </p>
    </footer>
</div>

<style lang="less">
    @import 'svelte-material-ui/bare.css';
    /*  */
    h1,
    .image,
    .button {
        text-align: center;
    }
    h1 {
        margin: 0;
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        font-weight: 400;
        font-size: 1.5rem;
        line-height: 1.334;
        letter-spacing: 0em;
        margin-bottom: 24px;
    }
    .container {
        margin: 68px auto;
        width: 396px;
    }

    .username,
    .rememberme,
    .button {
        margin-bottom: 20px;
    }

    .rememberme {
        margin-left: -8px;
        margin-top: 5px;
        margin-bottom: 20px;
    }
    .others {
        display: flex;
        justify-content: space-between;
        font-size: 0.875rem;
    }
    a {
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        color: #1976d2;
        text-decoration: underline;
        text-decoration-color: rgba(25, 118, 210, 0.4);
    }
    a:hover {
        text-decoration-color: inherit;
    }
    /* 输入框 */
    :global(.shaped-outlined) {
        width: 100%;
    }
    /* :global(.mdc-button__ripple) {
        width: 100%;
    } */
    :global(.mdc-button) {
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        font-weight: 500;
        font-size: 0.875rem;
        color: #fff;
        background-color: #1976d2;
        width: 100%;
    }
    :global(.mdc-button:not(:disabled)) {
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        font-weight: 500;
        font-size: 0.875rem;
        color: #fff;
        background-color: #1976d2;
        width: 100%;
    }

    footer {
        margin-top: 64px;
        margin-bottom: 32px;
        text-align: center;
        color: rgba(0, 0, 0, 0.6);
        font-size: 0.875rem;
        font-family: "Roboto", "Helvetica", "Arial", sans-serif;
        bottom: 0px;
    }
</style>
