package root

// Body html page content
const Body = `
<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="hostname/upload" method="post">
    <input type="file" name="uploadfile"/>
    <input type="text" name="phone"/>
    <input type="checkbox" name="private" value="private"/>
    <input type="hidden" name="MD5" value="1bc29b36f623ba82aaf6724fd3b16718"/>
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload"/>
</form>
</body>
</html>
`
