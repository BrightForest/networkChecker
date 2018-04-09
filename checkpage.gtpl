<html>
    <head>
    <title>Network Checker</title>
    </head>
    <body>
        <form action="/" method="post">
            IP Address:<input type="text" name="ip">
            <br>Port:<input type="text" name="port"></br>
            <br><input type="submit" value="Check!"></br>
            <br>{{.Answer}}</br>
        </form>
    </body>
</html>