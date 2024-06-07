package main

import (
	"fmt"
    "math/rand"
    "time"
)

// 방패 (방어구)
// 마나석 강화
var nickname string
var Shield string = "아이언우드 방패"
var EquipLev int = 15 // 레벨 15 장비
var Guard int = 10 // 방어력 +10
var DR float32 = 2.5 // 피해 감소 2.5% Damage Reduction
var Upgrade int = 0 // 0회 강화
var TryEnchant string
var SuccessChance int // 성공 확률
var WizCoin int = 0 // 소모된 위즈
var ManaStone int = 0 // 마나석
var TotalManaStone int = 0 // 소모된 마나석
var RadiantStone int = 0 // 소모된 광휘석
var RadiantOption int = 0 // 광휘석 옵션
var RetryEnchant string

func main() {
	name()
}

func name() {
	fmt.Print("플레이어의 이름을 입력하세요 : ")
	fmt.Scan(&nickname)

    fmt.Printf("안녕하세요! %s 님!\n\n", nickname)
        
    Enchant()
}

func Enchant() {
    if (Upgrade < 5) { // 업그레이드 최대 5회
        for {
            if Upgrade > 0 {
                fmt.Printf("장비 : %s (Lv.%d) +%d\n", Shield, EquipLev, Upgrade)
            } else {
                fmt.Printf("장비 : %s (Lv.%d)\n", Shield, EquipLev)
            }
            fmt.Printf("방어력 +%d / 피해 감소 %.2f\n\n", Guard, DR)
            fmt.Printf("소모된 위즈 코인 : %d 개\n\n", WizCoin)
            fmt.Print("마나석 몇 개를 사용하시겠습니까? (1 ~ 3) : ")
            fmt.Scan(&ManaStone)
            fmt.Printf("\n")
            
            switch ManaStone {
            case 1:
                SuccessChance = 30
                WizCoin += (ManaStone * 150)
                fmt.Printf("강화 성공 확률 : %d %%\n", SuccessChance)
                fmt.Printf("위즈 코인 소모량 : %d 개\n", (ManaStone * 150))
                Enchanting()
                return
            case 2:
                SuccessChance = 50
                WizCoin += (ManaStone * 150)
                fmt.Printf("강화 성공 확률 : %d %%\n", SuccessChance)
                fmt.Printf("위즈 코인 소모량 : %d 개\n", (ManaStone * 150))
                Enchanting()
                return
            case 3:
                SuccessChance = 70
                WizCoin += (ManaStone * 150)
                fmt.Printf("강화 성공 확률 : %d %%\n", SuccessChance)
                fmt.Printf("위즈 코인 소모량 : %d 개\n", (ManaStone * 150))
                Enchanting()
                return
            default:
                fmt.Printf("잘못 입력하셨습니다. 마력석 사용 가능 갯수는 최대 3개입니다.\n")
                continue
            }
        }
    } else if (Upgrade == 5) {
        if RadiantOption == 0 {
            fmt.Printf("장비 : %s (Lv.%d) +%d\n", Shield, EquipLev, Upgrade)
            fmt.Printf("방어력 +%d / 피해 감소 %.2f\n\n", Guard, DR)
            fmt.Printf("소모된 위즈 코인 : %d 개\n", WizCoin)
            RadiantEnchant()
        } else {
            fmt.Printf("장비 : %s (Lv.%d) +%d\n", Shield, EquipLev, Upgrade)
            fmt.Printf("방어력 +%d / 피해 감소 %.2f\n", Guard, DR)
            if RadiantOption == 1 {
                fmt.Printf("추가옵션 : [방어력 +10]\n\n")
            } else if RadiantOption == 2 {
                fmt.Printf("추가옵션 : [피해 감소 +2%%]\n\n")
            } else if RadiantOption == 3 {
                fmt.Printf("추가옵션 : [MaxHP +400]\n\n")
            } else if RadiantOption == 4 {
                fmt.Printf("추가옵션 : [MaxMP +150]\n\n")
            } else if RadiantOption == 5 {
                fmt.Printf("추가옵션 : [물리 속성 내성 +4%%]\n\n")
            } else if RadiantOption == 6 {
                fmt.Printf("추가옵션 : [방어력 +25]\n\n")
            } else if RadiantOption == 7 {
                fmt.Printf("추가옵션 : [피해 감소 +1%% / 물리 속성 내성 +3%%]\n\n")
            } else if RadiantOption == 8 {
                fmt.Printf("추가옵션 : [방어력 +20 / MaxHP +240]\n\n")
            }
            fmt.Printf("소모된 위즈 코인 : %d 개\n", WizCoin)
            fmt.Printf("광휘석 강화를 재시도 하겠습니까? (y / n) : ")
            fmt.Scan(&RetryEnchant)
            fmt.Printf("\n")
            
            switch RetryEnchant {
            case "y":
                RadiantEnchant()
            case "n":
                fmt.Printf("장비 아이템 강화를 종료합니다... 수고하셨습니다! %s 님!\n", nickname)
            default:
                fmt.Printf("잘못된 입력입니다...\n\n")
                Enchant()
            }
        }
    }
}

func Enchanting() {

    rand.Seed(time.Now().UnixNano()) // 시드 설정
    randomSuccess := rand.Intn(100) // 0부터 99 사이의 난수 생성
    randomIncrease := (rand.Intn(4) + 1) // 1부터 5 사이의 난수 생성 (방어력 증가)

	fmt.Printf("사용할 마나석 %d개\n", ManaStone)
	fmt.Printf("강화 하시겠습니까? (y / n) : ")
	fmt.Scan(&TryEnchant)
	fmt.Printf("\n")

    switch TryEnchant {
    case "y":
		if randomSuccess < SuccessChance {
			Upgrade++
			Guard += randomIncrease
			DR += 0.1
			fmt.Printf("강화에 성공하였습니다!\n")
			fmt.Printf("방어력이 %d 증가하였습니다.\n", randomIncrease)
			fmt.Printf("피해 감소율이 0.1 증가하였습니다.\n\n")
		} else {
			fmt.Printf("강화에 실패하였습니다...\n\n")
		}
		Enchant()
    case "n":
        WizCoin -= (ManaStone * 150)
        fmt.Printf("마나석 강화를 취소합니다.\n\n")
        Enchant()
    default:
        WizCoin -= (ManaStone * 150)
        fmt.Printf("잘못된 입력입니다...\n\n")
        Enchant()
    }
}

func RadiantEnchant() {

    rand.Seed(time.Now().UnixNano()) // 시드 설정
    radiantSuccess := rand.Intn(100) // 0부터 99 사이의 난수 생성
    randomOption := (rand.Intn(8) + 1) // 옵션 경우의 수 8가지

    fmt.Printf("장비 아이템에 광휘석을 사용하시겠습니까? (y / n) : ")
    fmt.Scan(&TryEnchant)
    fmt.Printf("\n")
    fmt.Printf("소모된 광휘석 : %d 개\n\n", RadiantStone)

    switch TryEnchant {
    case "y":
        RadiantStone++
        if radiantSuccess < 60 {
            fmt.Printf("★ 광휘석 강화에 성공했습니다!! ★\n")
            if randomOption == 1 {
                fmt.Printf("추가옵션 : [방어력 +10]\n\n")
                RadiantOption = 1
            } else if randomOption == 2 {
                fmt.Printf("추가옵션 : [피해 감소 +2%%]\n\n")
                RadiantOption = 2
            } else if randomOption == 3 {
                fmt.Printf("추가옵션 : [MaxHP +400]\n\n")
                RadiantOption = 3
            } else if randomOption == 4 {
                fmt.Printf("추가옵션 : [MaxMP +150]\n\n")
                RadiantOption = 4
            } else if randomOption == 5 {
                fmt.Printf("추가옵션 : [물리 속성 내성 +4%%]\n\n")
                RadiantOption = 5
            } else if randomOption == 6 {
                fmt.Printf("추가옵션 : [방어력 +25]\n\n")
                RadiantOption = 6
            } else if randomOption == 7 {
                fmt.Printf("추가옵션 : [피해 감소 +1%% / 물리 속성 내성 +3%%]\n\n")
                RadiantOption = 7
            } else if randomOption == 8 {
                fmt.Printf("추가옵션 : [방어력 +20 / MaxHP +240]\n\n")
                RadiantOption = 8
            }
        } else {
			fmt.Printf("광휘석 강화에 실패하였습니다...\n\n")
		}
        Enchant()
    case "n":
        fmt.Printf("광휘석 강화를 취소합니다.\n\n")
        Enchant()
    default:
        fmt.Printf("잘못된 입력입니다...\n\n")
        Enchant()
    }
}