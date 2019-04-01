package main

import (
  flatbuffers "github.com/google/flatbuffers/go"
  sample "github.com/jayrbolton/flatbuffer_examples/MyGame/Sample"
  "io/ioutil"
  "log"
)

func main () {
  BuildOrc()
  ReadOrc()
}

func BuildOrc () {

  builder := flatbuffers.NewBuilder(1024)
  weaponOne := builder.CreateString("Sword")
  weaponTwo := builder.CreateString("Axe")

  // Create the first weapon
  sample.WeaponStart(builder)
  sample.WeaponAddName(builder, weaponOne)
  sample.WeaponAddDamage(builder, 3)
  sword := sample.WeaponEnd(builder)

  // Second weapon
  sample.WeaponStart(builder)
  sample.WeaponAddName(builder, weaponTwo)
  sample.WeaponAddDamage(builder, 5)
  axe := sample.WeaponEnd(builder)

  name := builder.CreateString("Orc")
  sample.MonsterStartInventoryVector(builder, 10)
  for i := 9; i >= 0; i-- {
    builder.PrependByte(byte(i))
  }
  inv := builder.EndVector(10)

  // Build the weapons vector
  sample.MonsterStartWeaponsVector(builder, 2)
  builder.PrependUOffsetT(axe)
  builder.PrependUOffsetT(sword)
  weapons := builder.EndVector(2)

  // Create path vector
  sample.MonsterStartPathVector(builder, 2)
  sample.CreateVec3(builder, 1.0, 2.0, 3.0)
  sample.CreateVec3(builder, 4.0, 5.0, 6.0)
  path := builder.EndVector(2)

  // Build the monster
  sample.MonsterStart(builder)
  sample.MonsterAddPos(builder, sample.CreateVec3(builder, 1.0, 2.0, 3.0))
  sample.MonsterAddHp(builder, 300)
  sample.MonsterAddName(builder, name)
  sample.MonsterAddInventory(builder, inv)
  sample.MonsterAddColor(builder, sample.ColorRed)
  sample.MonsterAddWeapons(builder, weapons)
  sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
  sample.MonsterAddEquipped(builder, axe)
  sample.MonsterAddPath(builder, path)
  orc := sample.MonsterEnd(builder)
  builder.Finish(orc)

  buf := builder.FinishedBytes()
  err := ioutil.WriteFile("./orc.bin", buf, 0644)
  if err != nil { panic(err) }
}

func ReadOrc () {
  dat, err := ioutil.ReadFile("./orc.bin")
  if err != nil { panic(err) }
  monster := sample.GetRootAsMonster(dat, 0)
  hp := monster.Hp()
  mana := monster.Mana()
  name := string(monster.Name())
  log.Printf("HP: %v\n", hp)
  log.Printf("Mana: %v\n", mana)
  log.Printf("Name: %v\n", name)

  // Access sub-objects
  pos := monster.Pos(nil)
  x := pos.X()
  y := pos.Y()
  z := pos.Z()
  log.Printf("Position: %v, %v, %v\n", x, y, z)

  // Index vectors
  invLength := monster.InventoryLength()
  thirdItem := monster.Inventory(2)
  log.Printf("Inventory length: %v\n", invLength)
  log.Printf("Third item: %v\n", thirdItem)

  // Show weapons (vector of tables)
  weapon := new(sample.Weapon)
  weaponsLen := monster.WeaponsLength()
  for i := 0; i < weaponsLen; i++ {
    if monster.Weapons(weapon, i) {
      log.Printf("Has weapon named %v\n", string(weapon.Name()))
      log.Printf("  with damage %v\n", weapon.Damage())
    }
  }

  // Show equipment
  unionTable := new(flatbuffers.Table)
  if monster.Equipped(unionTable) {
    unionType := monster.EquippedType()
    if unionType == sample.EquipmentWeapon {
      unionWeapon := new(sample.Weapon)
      unionWeapon.Init(unionTable.Bytes, unionTable.Pos)
      log.Printf("Currently has %v equipped\n", string(unionWeapon.Name()))
    }
  }
}
