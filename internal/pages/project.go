package pages

import (
	// "fmt"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func ProjectPage(app *tview.Application) *tview.Grid {
	// ============ DESINIG LOGIC ==============
	// header and its container
	header := tview.NewTextView().
		SetText("Projetos").
		SetTextAlign(tview.AlignCenter).
		SetTextStyle(tcell.StyleDefault.Bold(true))

	headerWrapper := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).      
		AddItem(header, 1, 0, true).    
		AddItem(nil, 0, 1, false)       





	// footer text and its container
	footerText := "(a) add  |  (d) delete  |  (e) edit   |  (m) move item"
	footer := tview.NewTextView().
        SetTextAlign(tview.AlignLeft).   
        SetText(footerText)

	

	// list 
	list := tview.NewList().
		AddItem("Item 1", "", '*', nil).
		AddItem("item 2", "", '*', nil).
		AddItem("item 3", "", '*', nil).
		AddItem("item 4", "", '*', nil).
		AddItem("item 5", "", '*', nil).
		AddItem("item 6", "", '*', nil)



	// description box
	descText := tview.NewTextView().
		SetText(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean semper ultricies eleifend. Proin auctor metus rhoncus felis hendrerit blandit. Integer maximus metus vitae ante semper, non egestas mi gravida. Aliquam leo ligula, fermentum a metus sollicitudin, malesuada dapibus nunc. Donec et placerat tortor, id tristique dolor. Pellentesque arcu felis, vestibulum nec elementum placerat, pulvinar vitae lacus. Mauris ut lacinia augue. Duis convallis non ex non gravida.

Fusce in neque eget mauris tempus laoreet eu nec neque. Nam fermentum mollis dolor, sit amet iaculis mauris consectetur vel. Pellentesque eget placerat libero. Ut consectetur justo mauris, sit amet rhoncus neque tincidunt ut. Fusce nec dui in eros eleifend ullamcorper et mattis felis. Pellentesque sodales purus ornare justo pellentesque, ut egestas eros tristique. Phasellus eget diam lacus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nulla egestas diam in tellus rhoncus, a vulputate augue commodo. Praesent auctor dignissim felis, rhoncus ultricies nibh feugiat eget. Integer euismod erat justo, non pellentesque metus porta bibendum. Nunc congue nisl in erat feugiat tincidunt. Curabitur placerat sapien ac arcu congue, a bibendum ante bibendum. Nam pulvinar eu erat ac blandit.

Ut in aliquam orci. Proin in sem nec erat sodales elementum. Phasellus porttitor mollis sapien, quis pulvinar orci efficitur et. Donec suscipit in sapien accumsan tempor. Ut tempus magna felis, id suscipit lorem consequat sit amet. Fusce eget mi scelerisque, elementum diam nec, euismod odio. Duis sit amet mollis mi. Nulla id bibendum arcu. Maecenas condimentum augue vel mi consequat, vitae consectetur ante porta. Maecenas scelerisque arcu sapien, at sollicitudin mi condimentum et. Nulla rutrum euismod massa, eu scelerisque enim varius quis. Suspendisse tincidunt lorem vel magna elementum facilisis. Morbi metus mauris, consequat non facilisis egestas, egestas a tellus.

Cras quis diam a dolor commodo consequat eu eu est. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Vestibulum accumsan orci sit amet turpis pharetra efficitur. Proin tincidunt, eros eget efficitur sodales, felis odio dignissim metus, vitae ornare erat neque at nulla. Pellentesque ut dapibus dui. Nam ut risus est. Cras id felis enim. In venenatis est nisl, a ultricies odio accumsan sed. Sed orci ex, tempus sit amet mauris vitae, viverra fringilla libero. Cras porttitor mauris turpis, eu pharetra purus scelerisque et. Sed dignissim lectus ut ultricies euismod.

Pellentesque bibendum elementum posuere. Suspendisse interdum, turpis ut sodales ultricies, ante magna gravida nulla, et mattis justo turpis id nibh. Mauris consequat ante quis sem semper bibendum. In hac habitasse platea dictumst. Donec libero felis, porta a massa eget, sagittis laoreet elit. Interdum et malesuada fames ac ante ipsum primis in faucibus. Vivamus vitae eros accumsan, tincidunt enim semper, vulputate lacus. Pellentesque faucibus nibh vitae commodo tincidunt.

Pellentesque id libero eu urna rhoncus ullamcorper. Morbi euismod egestas tortor id egestas. Morbi varius, tortor vitae tincidunt mattis, odio nulla gravida neque, at porttitor ante quam at ligula. Nulla bibendum egestas elit, quis sagittis eros bibendum non. Quisque mattis luctus nisl quis fringilla. Quisque elit magna, rutrum vitae condimentum sed, ullamcorper fermentum eros. In ultrices pharetra ligula, malesuada facilisis felis auctor quis. Praesent aliquet, tellus nec mattis vulputate, diam lectus suscipit arcu, nec finibus ligula ligula ac mauris. Etiam egestas, ipsum a efficitur malesuada, turpis erat luctus nisi, et blandit enim quam vitae libero. Quisque id commodo ante. Sed tempus magna eget nisl maximus, non mattis mi malesuada. Pellentesque at sodales metus, at iaculis orci. Vestibulum rhoncus nisi quam. Nullam justo mi, facilisis sit amet convallis sed, dictum ut tellus. Vivamus vitae lacus mauris.

Integer at luctus sem. Praesent ut mauris et nisl efficitur iaculis. Nullam rutrum est eu eros dignissim, eget congue urna sollicitudin. Nam varius aliquet suscipit. Duis vehicula diam ut metus mattis vulputate. Donec at magna neque. Mauris purus lectus, rutrum id ex quis, auctor pretium arcu.

Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Suspendisse a nisi congue, mollis sem a, cursus arcu. Cras id laoreet eros. Sed ut interdum urna. Morbi interdum pulvinar libero, vitae gravida justo semper id. Ut vel viverra tellus. Morbi erat purus, ornare laoreet ipsum nec, pretium dictum arcu. Nulla a quam eu lectus dignissim sodales pharetra in risus. Aliquam dictum feugiat dolor eget congue. Nulla dolor nibh, bibendum quis bibendum sed, feugiat sed leo. Ut venenatis turpis nunc, eget sagittis enim aliquet eget. Praesent dolor arcu, mollis nec ipsum nec, malesuada imperdiet elit. Integer dictum mollis purus fermentum euismod. Interdum et malesuada fames ac ante ipsum primis in faucibus.

Aliquam eu rhoncus risus, sit amet faucibus lorem. Aliquam erat volutpat. Aliquam pulvinar lacus eget massa tristique laoreet. Pellentesque et vestibulum risus. Suspendisse potenti. Duis nibh metus, egestas vitae posuere vehicula, blandit ut ex. Pellentesque nibh lectus, ornare nec mollis ut, tempus nec felis. Vivamus sit amet facilisis purus. Aenean tempor aliquam faucibus. Ut et egestas ligula. Integer lobortis orci urna, vel sollicitudin libero eleifend sit amet. Quisque malesuada, nisl viverra ultricies iaculis, elit nisl interdum nibh, et tristique justo quam sed quam. Praesent maximus aliquet sapien, vitae venenatis nisi accumsan imperdiet.

Cras in interdum tellus. Ut imperdiet mauris nec ligula pulvinar finibus. Quisque aliquam tristique ex, nec convallis libero malesuada eu. Nullam facilisis, risus posuere convallis egestas, augue elit suscipit turpis, nec elementum odio leo nec dui. Curabitur dui dolor, semper vitae lectus ac, vehicula rhoncus elit. Duis ut urna non nibh laoreet condimentum sed et neque. Nunc a consequat arcu. Proin sed iaculis sem, vitae vehicula ipsum. Vivamus hendrerit dignissim nisl, nec tincidunt turpis ultricies vel. Nam eget euismod eros. Integer a rutrum augue. Vivamus mauris tortor, imperdiet vitae ligula eget, tincidunt malesuada erat. Praesent vitae facilisis velit. Phasellus ultricies nulla vel metus auctor hendrerit. Nunc malesuada eget purus vel maximus.

Aliquam nec risus sit amet tortor tempor scelerisque. Nullam fermentum sapien et viverra sollicitudin. Quisque lobortis imperdiet tortor. Curabitur a pharetra ipsum, eu ornare est. Praesent id risus sed magna tempus feugiat eu a mi. Quisque a nulla a augue efficitur laoreet. Integer volutpat, nunc sit amet sagittis iaculis, urna nibh semper nunc, nec mollis purus arcu quis massa. Aenean fringilla ipsum non dolor lobortis, ut congue metus blandit. Phasellus rhoncus ante id eros molestie consequat.

Donec sodales metus lorem, id venenatis purus ultricies sed. Nulla sit amet nunc at nisl viverra mollis. Maecenas at massa maximus, scelerisque est eget, porttitor dui. Sed facilisis dolor varius egestas consectetur. Cras lobortis ultrices mattis. Etiam odio dolor, finibus quis velit sit amet, aliquet tincidunt velit. Aenean hendrerit dapibus nisi, vitae condimentum massa interdum vel. Nam vitae ultrices nibh. Vestibulum fringilla bibendum luctus. Vivamus id nulla efficitur nibh fermentum hendrerit sed id ligula. Ut vitae tincidunt mauris. Nam at ipsum congue, varius nisi id, fringilla nisi.

Morbi malesuada purus sit amet tristique tincidunt. Pellentesque sit amet ornare magna. Mauris eget sapien euismod nibh tristique mollis. Vivamus molestie tortor nec sapien lacinia, ut dapibus leo vulputate. Sed ornare in massa id volutpat. Praesent at gravida ligula, quis tristique augue. Curabitur ullamcorper vehicula odio, quis porttitor lectus pulvinar vel. Phasellus viverra lacus dapibus dapibus placerat. Nulla ut augue orci. Nulla facilisi. Proin dignissim turpis facilisis, venenatis dolor sed, commodo magna. Proin placerat elit erat, ut tempor quam suscipit a. Nam gravida justo a tempor ultricies. Cras eget dolor cursus dolor aliquet auctor blandit eu dolor.

Etiam interdum justo tortor, at ultrices lacus dictum sit amet. Nullam magna magna, vulputate vitae sem eget, ultricies imperdiet dolor. Vestibulum nec erat vel risus hendrerit auctor vitae ut leo. Integer eget lobortis quam. Vivamus vehicula, quam a mattis cursus, nibh sapien finibus risus, molestie rhoncus felis magna vitae sapien. Sed purus magna, lacinia sed pretium id, congue id felis. Phasellus finibus, libero quis tincidunt aliquet, eros lectus ornare ex, sit amet bibendum ipsum ligula eget urna. Nullam auctor enim eu tellus lobortis cursus. Praesent mi odio, venenatis eu ante eget, viverra imperdiet dolor. Nam quis turpis metus. Sed egestas placerat interdum. Donec ullamcorper sagittis cursus. Suspendisse sodales magna et consectetur luctus.

Morbi sed tincidunt lacus. Vestibulum et enim neque. Pellentesque iaculis turpis ut risus tincidunt, eu gravida dolor fringilla. Aliquam sit amet convallis tortor. Vivamus dignissim nulla lorem, tempor euismod orci gravida a. Nulla facilisi. Cras vel nibh a lectus elementum lacinia sed a erat. Curabitur dignissim, nunc ut imperdiet accumsan, urna enim ullamcorper libero, et sodales massa magna at leo. Duis auctor ullamcorper elementum. Suspendisse potenti.

Integer sit amet lorem eget ipsum tristique sagittis nec eget neque. Duis nec nulla non orci sodales venenatis. In leo massa, mattis sagittis ipsum ac, tempus facilisis ante. Vivamus nibh odio, vestibulum vel faucibus vitae, porttitor porttitor ante. Suspendisse ultrices gravida faucibus. Proin semper commodo velit, vitae efficitur erat tempor sit amet. Vestibulum ultricies, est sit amet vehicula fermentum, purus lorem porttitor leo, et egestas ipsum lectus non eros. Nunc feugiat eu orci non pulvinar. Donec fringilla elementum odio ut auctor. Sed placerat aliquet neque.

Pellentesque tempor, turpis sit amet vehicula blandit, leo dolor varius nisi, nec ornare ipsum sem nec nulla. Vivamus vitae metus quis nisl laoreet sollicitudin non ac lorem. Nam sit amet magna lacus. Phasellus nisi dolor, consequat nec elementum a, accumsan eget metus. Phasellus hendrerit bibendum nunc, sed mattis sem porttitor eu. Praesent commodo nibh at sem tempor, vel fermentum odio varius. Nullam facilisis diam sem, in facilisis quam imperdiet ac. Vivamus commodo pellentesque dui, finibus suscipit nibh mattis non. In erat velit, blandit a eros in, faucibus egestas eros. Integer condimentum vehicula arcu eu suscipit. Suspendisse potenti. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aenean sit amet tortor et magna accumsan posuere nec quis massa.

Aenean non finibus sapien. Cras nec fermentum nulla. Sed tempor nibh sed velit hendrerit porta. Maecenas rutrum, sem vel congue tempus, nunc dolor pharetra dui, quis consequat odio lacus eu ligula. Etiam elementum mauris nec justo consectetur tristique. Etiam nulla ligula, suscipit eget arcu in, mollis posuere odio. Aliquam vulputate augue in gravida rutrum. Aliquam vel ligula eget ligula aliquam ornare. Pellentesque pretium sodales rutrum. Ut ut porttitor nisl. Cras tempor fermentum interdum.

Integer a nisl lacus. Proin at felis erat. Integer et fermentum sapien. Phasellus ex leo, sollicitudin quis erat ac, luctus dignissim leo. Suspendisse mollis augue sagittis, mattis nulla ut, faucibus leo. Vivamus semper, neque in feugiat ullamcorper, mi sem ullamcorper nulla, vitae commodo erat erat vitae turpis. Pellentesque sed diam non tellus aliquam volutpat. Donec iaculis mattis ante. Aliquam et velit sed nisi vulputate ullamcorper. Morbi vestibulum id quam at faucibus. Vestibulum quam nibh, egestas sed mi in, rutrum laoreet sapien. Phasellus et ante erat. Curabitur vel massa eu urna eleifend scelerisque.

Etiam venenatis dolor quis risus ornare dictum. In id ex a est sollicitudin gravida. Quisque risus mauris, cursus ac arcu at, finibus efficitur augue. Mauris et rhoncus tortor, non lobortis dolor. Sed luctus consequat ultrices. Nullam molestie augue eget lacus vulputate pretium. Aliquam eleifend lorem nec dolor efficitur, eu tempus neque dictum. Vivamus porta pulvinar lorem, non consectetur urna fermentum posuere. Nam eu pretium turpis, eget faucibus neque. Interdum et malesuada fames ac ante ipsum primis in faucibus.

Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nullam volutpat blandit ipsum, eget finibus dui. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam rhoncus eros augue, eu fermentum lorem posuere sit amet. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent iaculis non nulla vitae pretium. Donec egestas tellus eu arcu viverra, non condimentum risus egestas. Nulla porta tortor et felis mattis maximus. Aliquam consequat nisl ut maximus blandit. Ut nisl orci, rutrum ultricies hendrerit ut, elementum vitae orci. Duis et quam nec metus vulputate ullamcorper. Fusce convallis lobortis arcu. Ut quis ligula a diam pulvinar dictum. Quisque iaculis ipsum ac sollicitudin facilisis. Mauris justo ipsum, fringilla sit amet commodo sollicitudin, maximus sit amet diam.

Morbi a enim lacus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Fusce viverra efficitur consectetur. Nullam iaculis odio vitae lacinia suscipit. Sed ac sapien augue. Nullam condimentum arcu diam, vitae dictum velit vestibulum nec. Praesent quam risus, accumsan eget nisl sit amet, consequat facilisis justo. In et risus vitae ante rhoncus molestie a in quam.

Sed commodo non nibh id dapibus. In auctor id felis non volutpat. Vivamus sed lorem vulputate, placerat ex ultricies, rutrum est. Proin elementum augue pretium, dictum enim nec, placerat eros. Etiam urna nisl, mattis non laoreet sed, luctus non massa. Morbi consectetur lectus at euismod posuere. Sed porttitor vehicula magna, ut condimentum leo viverra non. Suspendisse auctor eros ultricies, convallis justo quis, accumsan tellus. Integer commodo, erat ac vestibulum commodo, turpis mauris scelerisque ante, ut suscipit magna libero id nibh. Curabitur quis eros luctus, posuere lectus a, facilisis ligula. Maecenas tempor ultricies sapien vitae convallis. Nunc egestas lobortis dolor, vel hendrerit leo pellentesque eget.

Vestibulum euismod et enim ac dignissim. Etiam nec sagittis ligula. Proin ut efficitur justo. Nulla non nulla volutpat, varius massa non, semper orci. Pellentesque ut porta ligula. Fusce dapibus a metus eu porttitor. Duis blandit ultricies mi sed blandit. Quisque pellentesque elit eu dignissim hendrerit. Nulla facilisi. Suspendisse et massa vehicula, volutpat risus eu, ullamcorper orci. Fusce porta ante nulla, at congue nibh vehicula sit amet. Pellentesque nulla turpis, vulputate ut velit sit amet, iaculis ultricies est. Nulla feugiat ut ligula eu dapibus. Aliquam eu luctus augue.

Proin pharetra purus in justo fermentum, sed commodo odio hendrerit. Vestibulum imperdiet est vel urna posuere pharetra vitae vitae massa. Proin tempor nisl cursus ullamcorper feugiat. Sed quam nunc, sodales sed dignissim eu, feugiat at enim. Vivamus justo turpis, faucibus at lacus efficitur, iaculis placerat ipsum. Ut ut enim nisl. Maecenas diam ipsum, ullamcorper vel dui id, vulputate elementum magna. Aenean eu lobortis sem, eget pulvinar justo. Etiam imperdiet vulputate sollicitudin. Praesent massa dolor, feugiat ut fringilla a, dictum ac nunc. Aenean nibh dui, lobortis vel efficitur eu, consectetur sed ex. Nullam eu lorem massa. Donec leo turpis, blandit eget blandit at, pulvinar ac lorem.

Duis aliquam laoreet bibendum. Cras eros nibh, feugiat vel justo ac, tincidunt ultrices quam. Nulla ex nulla, mollis efficitur felis ut, rhoncus semper turpis. Maecenas iaculis elit nulla, sed gravida ligula convallis id. Proin in molestie risus, luctus bibendum mi. In commodo suscipit leo eget posuere. Phasellus elementum arcu in eros tristique auctor. In euismod, dolor id volutpat sollicitudin, est dui pretium lectus, sit amet bibendum arcu nisl at nunc. Fusce suscipit tellus nibh, eget vestibulum velit fringilla quis. In hac habitasse platea dictumst. Etiam tincidunt augue vitae egestas vestibulum. Etiam ut nibh in odio lacinia ultrices. Mauris tincidunt erat eget tempus eleifend.

In tempor nisi metus, in facilisis lacus rutrum eu. Curabitur eleifend ornare justo, id euismod nulla fermentum eu. Aliquam ullamcorper mattis est in cursus. Donec non fermentum nisl. Integer nunc libero, placerat sit amet lacus in, malesuada sagittis libero. Praesent a enim ligula. Praesent et dictum quam. Aliquam erat volutpat. Sed massa nibh, finibus eget lacus ac, lacinia tempus odio. Phasellus eget posuere elit. Quisque volutpat ante non libero condimentum commodo et sed nunc. Proin arcu mi, maximus commodo faucibus eu, aliquam vitae est. Nam vitae massa hendrerit, faucibus metus a, efficitur metus. Suspendisse vel dui in mi euismod ullamcorper. Nunc aliquet molestie tortor, vel dignissim erat feugiat a. Vivamus convallis mi felis, in vestibulum lectus aliquam sit amet.

Donec magna metus, faucibus at metus ac, consectetur sodales lorem. In eu sem ante. Fusce turpis dui, rutrum tristique semper ut, vulputate ut mi. Etiam et justo eget lectus egestas luctus porttitor sed turpis. Morbi ac neque vel ipsum sollicitudin scelerisque. Vivamus ac erat cursus, molestie velit vel, pretium turpis. In mollis metus non malesuada luctus.

Vivamus ac augue purus. Donec nec cursus tortor, eu sollicitudin felis. Nullam a sollicitudin magna. Praesent accumsan tempor ultricies. Sed tincidunt, ex quis congue pellentesque, diam enim tristique felis, quis venenatis eros mauris ac dolor. Aliquam laoreet lobortis sodales. Duis nec elit tortor. Mauris massa ipsum, rhoncus et lobortis et, pretium eget diam.

Nulla facilisi. Ut at leo eu purus interdum posuere. Vestibulum mollis congue felis vitae feugiat. Nunc libero urna, imperdiet vitae commodo eu, viverra ut ex. Mauris commodo lorem in dignissim bibendum. Ut purus leo, sagittis sed metus at, gravida consectetur diam. Vivamus eros massa, tristique sed convallis ac, maximus a dolor. Aenean in orci velit. Sed quis congue purus, sit amet euismod dolor. Vestibulum consequat hendrerit eros nec maximus. Ut metus purus, porta vitae faucibus in, malesuada posuere eros. Curabitur viverra justo diam, aliquam gravida erat interdum in. Duis nec molestie nulla, et porttitor leo. Praesent nec metus nec purus cursus commodo ut sit amet leo. Nullam mollis, erat nec posuere semper, augue dolor aliquet justo, sit amet placerat nisi libero a metus. Proin ac congue nulla, sit amet egestas erat.`).
		SetWrap(true).
		SetScrollable(true)



	// Description
	description := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
		tview.NewTextView().
			SetText("Description").
			SetTextAlign(tview.AlignCenter).
			SetTextStyle(tcell.StyleDefault.Bold(true)),
		1, 0, true,
		).
		AddItem(nil, 1, 0, false). 
		AddItem(descText, 0, 1, true)


	// ============ KEYS LOGIC ==============
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		// move down at list
		case 'j':
			current := list.GetCurrentItem()
			if current < list.GetItemCount()-1 {
				list.SetCurrentItem(current + 1)
			}
			return nil

		// move up at list
		case 'k':
			current := list.GetCurrentItem()
			if current > 0 {
				list.SetCurrentItem(current - 1)
			}
			return nil

		// move to the thirst item
		case 'g':
			list.SetCurrentItem(0)

		// move to the last item
		case 'G':
			lastIndex := list.GetItemCount() - 1
			list.SetCurrentItem(lastIndex)
		}

		switch event.Key() {
		// move description up
		case tcell.KeyUp:
			row, col := descText.GetScrollOffset()
			if row > 0 {
				descText.ScrollTo(row-1, col)
			}
			return nil

		// move description donw
		case tcell.KeyDown:
			row, col := descText.GetScrollOffset()
			descText.ScrollTo(row+1, col)
			return nil

		// move the descripion to the begin
		case tcell.KeyPgUp:
			_, col := descText.GetScrollOffset()
			descText.ScrollTo(0, col)
			return nil

		// move the description to the end
		case tcell.KeyPgDn:
			_, col := descText.GetScrollOffset()
			descText.ScrollTo(999999, col)
			return nil


		default:
			return event
		}

	})



	// ============ GRID LOGIC ==============
	layout := tview.NewGrid().
		SetRows(3, 0, 1).
		SetColumns(120, 0).
		SetBorders(true).
		AddItem(headerWrapper, 0, 0, 1, 3, 0, 0, false).
		AddItem(description, 1, 0, 1, 1, 0, 0, false).
		AddItem(list, 1, 1, 1, 2, 0, 0, true).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	return layout
}

