package snapshotsync

import (
	"errors"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/ledgerwatch/turbo-geth/params"
)

const (
	DefaultChunkSize = 1024 * 1024
	LmdbFilename     = "data.mdb"
	MdbxFilename     = "mdbx.dat"
	EpochSize = 500_000

	//todo It'll be changed after enabling new snapshot generation mechanism
	HeadersSnapshotHash  = "4dcebdf20f67ce0a478fd5059a4c613ac961e138"
	BlocksSnapshotHash   = "296f1703f68afb46c3df040b097e2628fc27a66d"
	StateSnapshotHash    = "8f024711b2c2c277109b44053fcaab1b13346e69"
	
	SnapshotInfoHashPrefix  = "ih"
	SnapshotInfoBytesPrefix = "ib"
)

var (
	TorrentHashes = map[uint64]map[SnapshotType]metainfo.Hash{
		params.MainnetChainConfig.ChainID.Uint64(): {
			SnapshotType_headers: metainfo.NewHashFromHex(HeadersSnapshotHash),
			SnapshotType_bodies:  metainfo.NewHashFromHex(BlocksSnapshotHash),
			SnapshotType_state:   metainfo.NewHashFromHex(StateSnapshotHash),
		},
	}
	ErrInvalidSnapshot = errors.New("this snapshot for this chainID not supported ")
)

var Trackers = [][]string{{
	"http://35.189.110.210:80/announce",
},{
	"udp://tracker.openbittorrent.com:80",
	"udp://tracker.openbittorrent.com:80",
	"udp://tracker.publicbt.com:80",
	"udp://coppersurfer.tk:6969/announce",
	"udp://open.demonii.com:1337",
	"http://bttracker.crunchbanglinux.org:6969/announce",
	"udp://wambo.club:1337/announce",
	"udp://tracker.dutchtracking.com:6969/announce",
	"udp://tc.animereactor.ru:8082/announce",
	"udp://tracker.justseed.it:1337/announce",
	"udp://tracker.leechers-paradise.org:6969/announce",
	"udp://tracker.opentrackr.org:1337/announce",
	"https://open.kickasstracker.com:443/announce",
	"udp://tracker.coppersurfer.tk:6969/announce",
	"udp://open.stealth.si:80/announce",
	"http://87.253.152.137/announce",
	"http://91.216.110.47/announce",
	"http://91.217.91.21:3218/announce",
	"http://91.218.230.81:6969/announce",
	"http://93.92.64.5/announce",
	"http://atrack.pow7.com/announce",
	"http://bt.henbt.com:2710/announce",
	"http://bt.pusacg.org:8080/announce",
	"https://tracker.bt-hash.com:443/announce",
	"udp://tracker.leechers-paradise.org:6969",
	"https://182.176.139.129:6969/announce",
	"udp://zephir.monocul.us:6969/announce",
	"https://tracker.dutchtracking.com:80/announce",
	"https://grifon.info:80/announce",
	"udp://tracker.kicks-ass.net:80/announce",
	"udp://p4p.arenabg.com:1337/announce",
	"udp://tracker.aletorrenty.pl:2710/announce",
	"udp://tracker.sktorrent.net:6969/announce",
	"udp://tracker.internetwarriors.net:1337/announce",
	"https://tracker.parrotsec.org:443/announce",
	"https://tracker.moxing.party:6969/announce",
	"https://tracker.ipv6tracker.ru:80/announce",
	"https://tracker.fastdownload.xyz:443/announce",
	"udp://open.stealth.si:80/announce",
	"https://gwp2-v19.rinet.ru:80/announce",
	"https://tr.kxmp.cf:80/announce",
	"https://explodie.org:6969/announce",
}, {
	"udp://zephir.monocul.us:6969/announce",
	"udp://tracker.torrent.eu.org:451/announce",
	"udp://tracker.uw0.xyz:6969/announce",
	"udp://tracker.cyberia.is:6969/announce",
	"http://tracker.files.fm:6969/announce",
	"udp://tracker.zum.bi:6969/announce",
	"http://tracker.nyap2p.com:8080/announce",
	"udp://opentracker.i2p.rocks:6969/announce",
	"udp://tracker.zerobytes.xyz:1337/announce",
	"https://tracker.tamersunion.org:443/announce",
	"https://w.wwwww.wtf:443/announce",
	"https://tracker.imgoingto.icu:443/announce",
	"udp://blokas.io:6969/announce",
	"udp://api.bitumconference.ru:6969/announce",
	"udp://discord.heihachi.pw:6969/announce",
	"udp://cutiegirl.ru:6969/announce",
	"udp://fe.dealclub.de:6969/announce",
	"udp://ln.mtahost.co:6969/announce",
	"udp://vibe.community:6969/announce",
	"http://vpn.flying-datacenter.de:6969/announce",
	"udp://eliastre100.fr:6969/announce",
	"udp://wassermann.online:6969/announce",
	"udp://retracker.local.msn-net.ru:6969/announce",
	"udp://chanchan.uchuu.co.uk:6969/announce",
	"udp://kanal-4.de:6969/announce",
	"udp://handrew.me:6969/announce",
	"udp://mail.realliferpg.de:6969/announce",
	"udp://bubu.mapfactor.com:6969/announce",
	"udp://mts.tvbit.co:6969/announce",
	"udp://6ahddutb1ucc3cp.ru:6969/announce",
	"udp://adminion.n-blade.ru:6969/announce",
	"udp://contra.sf.ca.us:6969/announce",
	"udp://61626c.net:6969/announce",
	"udp://benouworldtrip.fr:6969/announce",
	"udp://sd-161673.dedibox.fr:6969/announce",
	"udp://cdn-1.gamecoast.org:6969/announce",
	"udp://cdn-2.gamecoast.org:6969/announce",
	"udp://daveking.com:6969/announce",
	"udp://bms-hosxp.com:6969/announce",
	"udp://teamspeak.value-wolf.org:6969/announce",
	"udp://edu.uifr.ru:6969/announce",
	"udp://adm.category5.tv:6969/announce",
	"udp://code2chicken.nl:6969/announce",
	"udp://t1.leech.ie:1337/announce",
	"udp://forever-tracker.zooki.xyz:6969/announce",
	"udp://free-tracker.zooki.xyz:6969/announce",
	"udp://public.publictracker.xyz:6969/announce",
	"udp://public-tracker.zooki.xyz:6969/announce",
	"udp://vps2.avc.cx:7171/announce",
	"udp://tracker.fileparadise.in:1337/announce",
	"udp://tracker.skynetcloud.site:6969/announce",
	"udp://z.mercax.com:53/announce",
	"https://publictracker.pp.ua:443/announce",
	"udp://us-tracker.publictracker.xyz:6969/announce",
	"udp://open.stealth.si:80/announce",
	"http://tracker1.itzmx.com:8080/announce",
	"http://vps02.net.orel.ru:80/announce",
	"http://tracker.gbitt.info:80/announce",
	"http://tracker.bt4g.com:2095/announce",
	"https://tracker.nitrix.me:443/announce",
	"udp://aaa.army:8866/announce",
	"udp://tracker.vulnix.sh:6969/announce",
	"udp://engplus.ru:6969/announce",
	"udp://movies.zsw.ca:6969/announce",
	"udp://storage.groupees.com:6969/announce",
	"udp://nagios.tks.sumy.ua:80/announce",
	"udp://tracker.v6speed.org:6969/announce",
	"udp://47.ip-51-68-199.eu:6969/announce",
	"udp://aruacfilmes.com.br:6969/announce",
	"https://trakx.herokuapp.com:443/announce",
	"udp://inferno.demonoid.is:3391/announce",
	"udp://publictracker.xyz:6969/announce",
	"http://tracker2.itzmx.com:6961/announce",
	"http://tracker3.itzmx.com:6961/announce",
	"udp://retracker.akado-ural.ru:80/announce",
	"udp://tracker-udp.gbitt.info:80/announce",
	"http://h4.trakx.nibba.trade:80/announce",
	"udp://tracker.army:6969/announce",
	"http://tracker.anonwebz.xyz:8080/announce",
	"udp://tracker.shkinev.me:6969/announce",
	"http://0205.uptm.ch:6969/announce",
	"udp://tracker.zooki.xyz:6969/announce",
	"udp://forever.publictracker.xyz:6969/announce",
	"udp://tracker.moeking.me:6969/announce",
	"udp://ultra.zt.ua:6969/announce",
	"udp://tracker.publictracker.xyz:6969/announce",
	"udp://ipv4.tracker.harry.lu:80/announce",
	"udp://u.wwwww.wtf:1/announce",
	"udp://line-net.ru:6969/announce",
	"udp://dpiui.reedlan.com:6969/announce",
	"udp://tracker.zemoj.com:6969/announce",
	"udp://t3.leech.ie:1337/announce",
	"http://t.nyaatracker.com:80/announce",
	"udp://exodus.desync.com:6969/announce",
	"udp://valakas.rollo.dnsabr.com:2710/announce",
	"udp://tracker.ds.is:6969/announce",
	"udp://tracker.opentrackr.org:1337/announce",
	"udp://tracker0.ufibox.com:6969/announce",
	"https://tracker.hama3.net:443/announce",
	"udp://opentor.org:2710/announce",
	"udp://t2.leech.ie:1337/announce",
	"https://1337.abcvg.info:443/announce",
	"udp://git.vulnix.sh:6969/announce",
	"udp://retracker.lanta-net.ru:2710/announce",
	"udp://tracker.lelux.fi:6969/announce",
	"udp://bt1.archive.org:6969/announce",
	"udp://admin.videoenpoche.info:6969/announce",
	"udp://drumkitx.com:6969/announce",
	"udp://tracker.dler.org:6969/announce",
	"udp://koli.services:6969/announce",
	"udp://tracker.dyne.org:6969/announce",
	"http://torrenttracker.nwc.acsalaska.net:6969/announce",
	"udp://rutorrent.frontline-mod.com:6969/announce",
	"http://rt.tace.ru:80/announce",
	"udp://explodie.org:6969/announce",
}, {
	"udp://public.popcorn-tracker.org:6969/announce",
	"http://104.28.1.30:8080/announce",
	"http://104.28.16.69/announce",
	"http://107.150.14.110:6969/announce",
	"http://109.121.134.121:1337/announce",
	"http://114.55.113.60:6969/announce",
	"http://125.227.35.196:6969/announce",
	"http://128.199.70.66:5944/announce",
	"http://157.7.202.64:8080/announce",
	"http://158.69.146.212:7777/announce",
	"http://173.254.204.71:1096/announce",
	"http://178.175.143.27/announce",
	"http://178.33.73.26:2710/announce",
	"http://182.176.139.129:6969/announce",
	"http://185.5.97.139:8089/announce",
	"http://188.165.253.109:1337/announce",
	"http://194.106.216.222/announce",
	"http://195.123.209.37:1337/announce",
	"http://210.244.71.25:6969/announce",
	"http://210.244.71.26:6969/announce",
	"http://213.159.215.198:6970/announce",
	"http://213.163.67.56:1337/announce",
	"http://37.19.5.139:6969/announce",
	"http://37.19.5.155:6881/announce",
	"http://46.4.109.148:6969/announce",
	"http://5.79.249.77:6969/announce",
	"http://5.79.83.193:2710/announce",
	"http://51.254.244.161:6969/announce",
	"http://59.36.96.77:6969/announce",
	"http://74.82.52.209:6969/announce",
	"http://80.246.243.18:6969/announce",
	"http://81.200.2.231/announce",
	"http://85.17.19.180/announce",
	"http://87.248.186.252:8080/announce",
	"http://87.253.152.137/announce",
	"http://91.216.110.47/announce",
	"http://91.217.91.21:3218/announce",
	"http://91.218.230.81:6969/announce",
	"http://93.92.64.5/announce",
	"http://atrack.pow7.com/announce",
	"http://bt.henbt.com:2710/announce",
	"http://bt.pusacg.org:8080/announce",
	"http://bt2.careland.com.cn:6969/announce",
	"http://explodie.org:6969/announce",
	"http://mgtracker.org:2710/announce",
	"http://mgtracker.org:6969/announce",
	"http://open.acgtracker.com:1096/announce",
	"http://open.lolicon.eu:7777/announce",
	"http://open.touki.ru/announce.php",
	"http://p4p.arenabg.ch:1337/announce",
	"http://p4p.arenabg.com:1337/announce",
	"http://pow7.com:80/announce",
	"http://retracker.gorcomnet.ru/announce",
	"http://retracker.krs-ix.ru/announce",
	"http://retracker.krs-ix.ru:80/announce",
	"http://secure.pow7.com/announce",
	"http://t1.pow7.com/announce",
	"http://t2.pow7.com/announce",
	"http://thetracker.org:80/announce",
	"http://torrent.gresille.org/announce",
	"http://torrentsmd.com:8080/announce",
	"http://tracker.aletorrenty.pl:2710/announce",
	"http://tracker.baravik.org:6970/announce",
	"http://tracker.bittor.pw:1337/announce",
	"http://tracker.bittorrent.am/announce",
	"http://tracker.calculate.ru:6969/announce",
	"http://tracker.dler.org:6969/announce",
	"http://tracker.dutchtracking.com/announce",
	"http://tracker.dutchtracking.com:80/announce",
	"http://tracker.dutchtracking.nl/announce",
	"http://tracker.dutchtracking.nl:80/announce",
	"http://tracker.edoardocolombo.eu:6969/announce",
	"http://tracker.ex.ua/announce",
	"http://tracker.ex.ua:80/announce",
	"http://tracker.filetracker.pl:8089/announce",
	"http://tracker.flashtorrents.org:6969/announce",
	"http://tracker.grepler.com:6969/announce",
	"http://tracker.internetwarriors.net:1337/announce",
	"http://tracker.kicks-ass.net/announce",
	"http://tracker.kicks-ass.net:80/announce",
	"http://tracker.kuroy.me:5944/announce",
	"http://tracker.mg64.net:6881/announce",
	"http://tracker.opentrackr.org:1337/announce",
	"http://tracker.skyts.net:6969/announce",
	"http://tracker.tfile.me/announce",
	"http://tracker.tiny-vps.com:6969/announce",
	"http://tracker.tvunderground.org.ru:3218/announce",
	"http://tracker.yoshi210.com:6969/announc",
	"http://tracker1.wasabii.com.tw:6969/announce",
	"http://tracker2.itzmx.com:6961/announce",
	"http://tracker2.wasabii.com.tw:6969/announce",
	"http://www.wareztorrent.com/announce",
	"http://www.wareztorrent.com:80/announce",
	"https://104.28.17.69/announce",
	"https://www.wareztorrent.com/announce",
	"udp://107.150.14.110:6969/announce",
	"udp://109.121.134.121:1337/announce",
	"udp://114.55.113.60:6969/announce",
	"udp://128.199.70.66:5944/announce",
	"udp://151.80.120.114:2710/announce",
	"udp://168.235.67.63:6969/announce",
	"udp://178.33.73.26:2710/announce",
	"udp://182.176.139.129:6969/announce",
	"udp://185.5.97.139:8089/announce",
	"udp://185.86.149.205:1337/announce",
	"udp://188.165.253.109:1337/announce",
	"udp://191.101.229.236:1337/announce",
	"udp://194.106.216.222:80/announce",
	"udp://195.123.209.37:1337/announce",
	"udp://195.123.209.40:80/announce",
	"udp://208.67.16.113:8000/announce",
	"udp://213.163.67.56:1337/announce",
	"udp://37.19.5.155:2710/announce",
	"udp://46.4.109.148:6969/announce",
	"udp://5.79.249.77:6969/announce",
	"udp://5.79.83.193:6969/announce",
	"udp://51.254.244.161:6969/announce",
	"udp://62.138.0.158:6969/announce",
	"udp://62.212.85.66:2710/announce",
	"udp://74.82.52.209:6969/announce",
	"udp://85.17.19.180:80/announce",
	"udp://89.234.156.205:80/announce",
	"udp://9.rarbg.com:2710/announce",
	"udp://9.rarbg.me:2780/announce",
	"udp://9.rarbg.to:2730/announce",
	"udp://91.218.230.81:6969/announce",
	"udp://94.23.183.33:6969/announce",
	"udp://bt.xxx-tracker.com:2710/announce",
	"udp://eddie4.nl:6969/announce",
	"udp://explodie.org:6969/announce",
	"udp://mgtracker.org:2710/announce",
	"udp://open.stealth.si:80/announce",
	"udp://p4p.arenabg.com:1337/announce",
	"udp://shadowshq.eddie4.nl:6969/announce",
	"udp://shadowshq.yi.org:6969/announce",
	"udp://torrent.gresille.org:80/announce",
	"udp://tracker.aletorrenty.pl:2710/announce",
	"udp://tracker.bittor.pw:1337/announce",
	"udp://tracker.coppersurfer.tk:6969/announce",
	"udp://tracker.eddie4.nl:6969/announce",
	"udp://tracker.ex.ua:80/announce",
	"udp://tracker.filetracker.pl:8089/announce",
	"udp://tracker.flashtorrents.org:6969/announce",
	"udp://tracker.grepler.com:6969/announce",
	"udp://tracker.ilibr.org:80/announce",
	"udp://tracker.internetwarriors.net:1337/announce",
	"udp://tracker.kicks-ass.net:80/announce",
	"udp://tracker.kuroy.me:5944/announce",
	"udp://tracker.leechers-paradise.org:6969/announce",
	"udp://tracker.mg64.net:2710/announce",
	"udp://tracker.mg64.net:6969/announce",
	"udp://tracker.opentrackr.org:1337/announce",
	"udp://tracker.piratepublic.com:1337/announce",
	"udp://tracker.sktorrent.net:6969/announce",
	"udp://tracker.skyts.net:6969/announce",
	"udp://tracker.tiny-vps.com:6969/announce",
	"udp://tracker.yoshi210.com:6969/announce",
	"udp://tracker2.indowebster.com:6969/announce",
	"udp://tracker4.piratux.com:6969/announce",
	"udp://zer0day.ch:1337/announce",
	"udp://zer0day.to:1337/announce",
}}
