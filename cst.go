package main

const templateHTML = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Result</title>
</head>

<style>
    body,
    html {
        background: black;
        font-weight: bold;
    }

    .five-star {
        background: rgb(71, 57, 62);
        color: #f0932b;
        padding: 2px;
        display: inline-block;
        margin-top: 2px;
    }

    .four-star {
        background: rgb(71, 57, 62);
        color: rgb(231, 152, 179);
        padding: 2px;
        display: inline-block;
        margin-top: 2px;
    }

    .sub-title {
        font-size: 12px;
    }
</style>

<body>
    <div style="background: #d35400; color: whitesmoke; padding: 10px;">
        角色活动祈愿 <span class="sub-title">总共{{ .CharacterGachaResult.Total }}抽 已{{ .CharacterGachaResult.Unluck
            }}抽未出金 平均{{ .CharacterGachaResult.AverageLuck }}抽出金</span>
    </div>
    <div style="margin-bottom: 20px;">
        {{ range $index, $luck := .CharacterGachaResult.LuckList }}
        {{if eq $luck.Type "5"}}
        <span class="five-star">[{{ $luck.Index }}]{{ $luck.Name }}</span><br>
        {{end}}
        {{if eq $luck.Type "4"}}
        <span class="four-star">[{{ $luck.Index }}]{{ $luck.Name }}</span>
        {{end}}
        {{ end }}
    </div>
    <div style="background: #16a085; color: whitesmoke; padding: 10px;">
        武器活动祈愿 <span class="sub-title">总共{{ .WeaponGachaResult.Total }}抽 已{{ .WeaponGachaResult.Unluck
            }}抽未出金 平均{{ .WeaponGachaResult.AverageLuck }}抽出金</span>
    </div>
    <div style="margin-bottom: 20px;">
        {{ range $index, $luck := .WeaponGachaResult.LuckList }}
        {{if eq $luck.Type "5"}}
        <span class="five-star">[{{ $luck.Index }}]{{ $luck.Name }}</span><br>
        {{end}}
        {{if eq $luck.Type "4"}}
        <span class="four-star">[{{ $luck.Index }}]{{ $luck.Name }}</span>
        {{end}}
        {{ end }}
    </div>
    <div style="background: #2980b9; color: whitesmoke; padding: 10px;">
        常驻祈愿 <span class="sub-title">总共{{ .OrdinaryGachaResult.Total }}抽 已{{ .OrdinaryGachaResult.Unluck
            }}抽未出金 平均{{ .OrdinaryGachaResult.AverageLuck }}抽出金</span>
    </div>
    <div style="margin-bottom: 20px;">
        {{ range $index, $luck := .OrdinaryGachaResult.LuckList }}
        {{if eq $luck.Type "5"}}
        <span class="five-star">[{{ $luck.Index }}]{{ $luck.Name }}</span><br>
        {{end}}
        {{if eq $luck.Type "4"}}
        <span class="four-star">[{{ $luck.Index }}]{{ $luck.Name }}</span>
        {{end}}
        {{ end }}
    </div>
</body>

</html>`
