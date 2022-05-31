function update_tokens(accounts) {
    if (!window.location.href.includes('snipers.html')) return;

    var ttl_snp = 0;

    document.getElementById('token_list').innerHTML = ""

    accounts.forEach(account => {
        ttl_snp += account.GuildSize;

        if (account.Avatar == '') {
            avatar_url = `https://external-preview.redd.it/4PE-nlL_PdMD5PrFNLnjurHQ1QKPnCvg368LTDnfM-M.png?auto=webp&s=ff4c3fbc1cce1a1856cff36b5d2a40a6d02cc1c3`
        } else {
            avatar_url = `https://cdn.discordapp.com/avatars/${account.ID}/${account.Avatar}`;
        }

        document.getElementById('token_list').innerHTML +=
            `
        <tr>
            <td class="p-2 whitespace-nowrap">
                <img class="w-10 h-10 rounded-full" src="${avatar_url}" alt="Avatar">
            </td>

            <td class="p-2 whitespace-nowrap">
                <div class="text-left font-medium text-gray-500">${account.Username}#${account.Discriminator}</div>
            </td>
            
            <td class="p-2 whitespace-nowrap">
                <div class="text-left font-medium text-green-500">${account.GuildSize}</div>
            </td>

            <td class="p-2 whitespace-nowrap">
                <div class="text-left font-medium text-orange-500">${account.RecievedMessages}</div>
            </td>

            <td class="p-2 whitespace-nowrap">
            <a href="#" class="font-medium text-gray-500 hover:underline" onclick='copy_token("${account.Token}")'>Get Token</a>
            </td>
        </tr>
        `
    });

    document.getElementById("sniper_count").innerHTML = `${accounts.length} Snipers monitoring ${ttl_snp} servers.`
};

function update_guilds(guilds) {
    if (!window.location.href.includes('guilds.html')) return;
    const urlParams = new URLSearchParams(window.location.search);

    let start_range = parseInt(urlParams.get('start'));
    let range = parseInt(urlParams.get('range'));

    document.getElementById('guild_list').innerHTML = ""

    for (var i = start_range; i < start_range + range; i++) {

        if (guilds[i].Avatar == "") {
            guilds[i].Avatar = 'https://polybit-apps.s3.amazonaws.com/stdlib/users/discord/profile/image.png?1621007833204'
        }

        document.getElementById('guild_list').innerHTML +=
            `
    <tr>
        <td class="p-2 whitespace-nowrap">
            <img class="w-10 h-10 rounded-full" src="${guilds[i].Avatar}" alt="Avatar">
        </td>

        <td class="p-2 whitespace-nowrap">
            <div class="text-left font-medium text-gray-500">${guilds[i].Name.substring(0,25)}</div>
        </td>
        
        <td class="p-2 whitespace-nowrap">
            <div class="text-left font-medium text-green-500">${guilds[i].Members}</div>
        </td>

        <td class="p-2 whitespace-nowrap">
            <div class="text-left font-medium text-orange-500">${guilds[i].Messages}</div>
        </td>

        <td class="p-2 whitespace-nowrap">
            <div class="text-left font-medium text-gray-500">${guilds[i].ID}</div>
        </td>
    </tr>
    `
    
    }

    document.getElementById("sniper_count").innerHTML = `${guilds.length} Servers.`

    document.getElementById("switcher").innerHTML = 
    `
    <center>
        <nav aria-label="Page navigation example">
            <ul class="inline-flex items-center -space-x-px">
                
                <li>
                    <a href="guilds.html?start=${start_range-20}&range=10"
                        class="py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">${start_range-20} - ${(start_range-20)+range}</a>
                </li>
                <li>
                    <a href="guilds.html?start=${start_range-10}&range=10"
                        class="py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">${start_range-10} - ${(start_range-10)+range}</a>
                </li>
                <li>
                    <a href="guilds.html?start=${start_range}&range=10" aria-current="page"
                        class="z-10 py-2 px-3 leading-tight text-blue-600 bg-blue-50 border border-blue-300 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white">${start_range}</a>
                </li>
                <li>
                    <a href="guilds.html?start=${start_range+10}&range=10"
                        class="py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">${start_range+10} - ${(start_range+10)+range}</a>
                </li>
                <li>
                    <a href="guilds.html?start=${start_range+20}&range=10"
                        class="py-2 px-3 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">${start_range+20} - ${(start_range+20)+range}</a>
                </li>
            </ul>
        </nav>
    </center>
    `
};

function copy_token(token) {
    var el = document.createElement('textarea');
    el.value = token;
    el.setAttribute('readonly', '');
    el.style = { position: 'absolute', left: '-9999px' };
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
}

window.onload = function () {
    let all_messages = 0;
    let tmp = 0;

    setInterval(async () => {
        if (!window.location.href.includes('index.html')) return;

        document.getElementById("mess_counter").innerHTML = `Total messages: ${all_messages}`;
        update_dstats(all_messages - tmp)
        tmp = all_messages;
    }, 1000)

    socket = new WebSocket("ws://proxies.gay:13254/ws");
    socket.onmessage = function (args) {
        const payload = JSON.parse(args.data);

        console.log(payload);

        switch (payload.Op) {
            case 'token_update':
                update_tokens(payload.AccountList);
                break;
            case 'update_dstat':
                all_messages = payload.TotalCheckedMessage;
                break;
            case 'guild_update':
                update_guilds(payload.GuildList)
        };
    };

    socket.onerror = function (error) {

    };
};