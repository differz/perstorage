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
    <input type="text" name="phone" placeholder="phone"/>
    <input type="text" name="description" placeholder="description"/>
    <input type="checkbox" name="private" value="private"/>
    <p><select size="6" name="category">
            <option disabled>Select category</option>
            <option value="1">Lifetime</option>
            <option selected value="2">Backup</option>
            <option value="4">Film</option>
            <option value="8">Music</option>
            <option value="16">Photo</option>
        </select>
    </p>
    <input type="hidden" name="MD5" value="1bc29b36f623ba82aaf6724fd3b16718"/>
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload"/>
</form>
</body>
</html>
`
