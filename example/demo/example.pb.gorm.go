// Code generated by protoc-gen-go-gorm. DO NOT EDIT.
// source: demo/example.proto

package example

import (
	json "encoding/json"
	pgtype "github.com/jackc/pgtype"
	pq "github.com/lib/pq"
	lo "github.com/samber/lo"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
)

type ThingGormModels []*ThingGormModel
type ThingProtos []*Thing

func (m ThingGormModels) ToProtos() (protos ThingProtos, err error) {
	protos = ThingProtos{}
	for _, model := range m {
		var proto *Thing
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p ThingProtos) ToModels() (models ThingGormModels, err error) {
	models = ThingGormModels{}
	for _, proto := range p {
		var model *ThingGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

type ThingGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{price:0.00,1000.00}"
	ADouble float64 `json:"aDouble" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	AFloat float32 `json:"aFloat" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{int32}"
	AnInt32 int32 `json:"anInt32" fake:"{int32}"`

	// @gotags: fake:"{number:9223372036854775807}"
	AnInt64 int64 `json:"anInt64" fake:"{number:9223372036854775807}"`

	// @gotags: fake:"{bool}"
	ABool bool `json:"aBool" fake:"{bool}"`

	// @gotags: fake:"{hackerphrase}"
	AString string `json:"aString" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	ABytes []byte `json:"aBytes" fake:"skip"`

	// @gotags: fake:"{price:0.00,1000.00}"
	Doubles pq.Float64Array `gorm:"type:float[];" json:"doubles" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	Floats pq.Float32Array `gorm:"type:float[];" json:"floats" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{int32}"
	Int32S pq.Int32Array `gorm:"type:int[];" json:"int32s" fake:"{int32}"`

	// @gotags: fake:"{number:9223372036854775807}"
	Int64S pq.Int64Array `gorm:"type:int[];" json:"int64s" fake:"{number:9223372036854775807}"`

	// @gotags: fake:"{bool}"
	Bools pq.BoolArray `gorm:"type:bool[];" json:"bools" fake:"{bool}"`

	// @gotags: fake:"{hackerphrase}"
	Strings pq.StringArray `gorm:"type:string[];" json:"strings" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	Bytess pq.ByteaArray `gorm:"type:bytes[];" json:"bytess" fake:"skip"`

	// @gotags: fake:"skip"
	OptionalScalarField *string `json:"optionalScalarField" fake:"skip"`

	// @gotags: fake:"skip"
	AStructpb *pgtype.JSONB `gorm:"type:jsonb" json:"aStructpb" fake:"skip"`

	BelongsToId *string ``

	// @gotags: fake:"skip"
	BelongsTo *BelongsToThingGormModel `gorm:"foreignKey:BelongsToId;" json:"belongsTo" fake:"skip"`

	// @gotags: fake:"skip"
	BelongsToTwoId string `json:"belongsToTwoId" fake:"skip"`

	// @gotags: fake:"skip"
	BelongsToTwo *BelongsToThingGormModel `gorm:"foreignKey:BelongsToTwoId;" json:"belongsToTwo" fake:"skip"`

	// @gotags: fake:"skip"
	AnUnexpectedId string `json:"anUnexpectedId" fake:"skip"`

	// @gotags: fake:"skip"
	BelongsToThree *BelongsToThingGormModel `gorm:"foreignKey:AnUnexpectedId;" json:"belongsToThree" fake:"skip"`

	// @gotags: fake:"skip"
	HasOne *HasOneThingGormModel `gorm:"foreignKey:ThingId;" json:"hasOne" fake:"skip"`

	// @gotags: fake:"skip"
	HasMany []*HasManyThingGormModel `gorm:"foreignKey:ThingId;" json:"hasMany" fake:"skip"`

	// @gotags: fake:"skip"
	ManyToMany []*ManyToManyThingGormModel `gorm:"many2many:things_many_to_many_things;" json:"manyToMany" fake:"skip"`

	// @gotags: fake:"{number:1,9}"
	IntEnum int `json:"intEnum" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	StringEnum string `json:"stringEnum" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	IntEnumList pq.Int32Array `gorm:"type:int[]" json:"intEnumList" fake:"{number:1,9}"`

	// @gotags: fake:"{number:1,9}"
	StringEnumList pq.StringArray `gorm:"type:string[]" json:"stringEnumList" fake:"{number:1,9}"`
}

func (m *ThingGormModel) TableName() string {
	return "things"
}

func (m *ThingGormModel) ToProto() (theProto *Thing, err error) {
	if m == nil {
		return
	}
	theProto = &Thing{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = m.UpdatedAt.Format(time.RFC3339Nano)
	}

	theProto.ADouble = m.ADouble

	theProto.AFloat = m.AFloat

	theProto.AnInt32 = m.AnInt32

	theProto.AnInt64 = m.AnInt64

	theProto.ABool = m.ABool

	theProto.AString = m.AString

	theProto.ABytes = m.ABytes

	theProto.Doubles = m.Doubles

	theProto.Floats = m.Floats

	theProto.Int32S = m.Int32S

	theProto.Int64S = m.Int64S

	theProto.Bools = m.Bools

	theProto.Strings = m.Strings

	theProto.Bytess = m.Bytess

	theProto.OptionalScalarField = m.OptionalScalarField

	if m.AStructpb != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(m.AStructpb); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theProto.AStructpb); err != nil {
			return
		}
	}

	if theProto.BelongsTo, err = m.BelongsTo.ToProto(); err != nil {
		return
	}

	theProto.BelongsToTwoId = m.BelongsToTwoId

	if theProto.BelongsToTwo, err = m.BelongsToTwo.ToProto(); err != nil {
		return
	}

	theProto.AnUnexpectedId = m.AnUnexpectedId

	if theProto.BelongsToThree, err = m.BelongsToThree.ToProto(); err != nil {
		return
	}

	if theProto.HasOne, err = m.HasOne.ToProto(); err != nil {
		return
	}

	if len(m.HasMany) > 0 {
		theProto.HasMany = []*HasManyThing{}
		for _, item := range m.HasMany {
			var HasManyProto *HasManyThing
			if HasManyProto, err = item.ToProto(); err != nil {
				return
			} else {
				theProto.HasMany = append(theProto.HasMany, HasManyProto)
			}
		}
	}

	if len(m.ManyToMany) > 0 {
		theProto.ManyToMany = []*ManyToManyThing{}
		for _, item := range m.ManyToMany {
			var ManyToManyProto *ManyToManyThing
			if ManyToManyProto, err = item.ToProto(); err != nil {
				return
			} else {
				theProto.ManyToMany = append(theProto.ManyToMany, ManyToManyProto)
			}
		}
	}

	theProto.IntEnum = EnumOne(m.IntEnum)

	theProto.StringEnum = EnumOne(EnumOne_value[m.StringEnum])

	if len(m.IntEnumList) > 0 {
		theProto.IntEnumList = []EnumOne{}
		for _, val := range m.IntEnumList {
			theProto.IntEnumList = append(theProto.IntEnumList, EnumOne(val))
		}
	}

	if len(m.StringEnumList) > 0 {
		theProto.StringEnumList = []EnumOne{}
		for _, val := range m.StringEnumList {
			theProto.StringEnumList = append(theProto.StringEnumList, EnumOne(EnumOne_value[val]))
		}
	}

	return
}

func (p *Thing) ToModel() (theModel *ThingGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &ThingGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != "" {
		var timestamp time.Time
		if timestamp, err = time.Parse(time.RFC3339Nano, p.UpdatedAt); err != nil {
			return
		}
		theModel.UpdatedAt = &timestamp
	}

	theModel.ADouble = p.ADouble

	theModel.AFloat = p.AFloat

	theModel.AnInt32 = p.AnInt32

	theModel.AnInt64 = p.AnInt64

	theModel.ABool = p.ABool

	theModel.AString = p.AString

	theModel.ABytes = p.ABytes

	theModel.Doubles = p.Doubles

	theModel.Floats = p.Floats

	theModel.Int32S = p.Int32S

	theModel.Int64S = p.Int64S

	theModel.Bools = p.Bools

	theModel.Strings = p.Strings

	theModel.Bytess = p.Bytess

	theModel.OptionalScalarField = p.OptionalScalarField

	if p.AStructpb != nil {
		var jsonBytes []byte
		if jsonBytes, err = json.Marshal(p.AStructpb); err != nil {
			return
		}
		if err = json.Unmarshal(jsonBytes, &theModel.AStructpb); err != nil {
			return
		}
	}

	if theModel.BelongsTo, err = p.BelongsTo.ToModel(); err != nil {
		return
	}

	theModel.BelongsToTwoId = p.BelongsToTwoId

	if theModel.BelongsToTwo, err = p.BelongsToTwo.ToModel(); err != nil {
		return
	}

	theModel.AnUnexpectedId = p.AnUnexpectedId

	if theModel.BelongsToThree, err = p.BelongsToThree.ToModel(); err != nil {
		return
	}

	if theModel.HasOne, err = p.HasOne.ToModel(); err != nil {
		return
	}

	if len(p.HasMany) > 0 {
		theModel.HasMany = []*HasManyThingGormModel{}
		for _, item := range p.HasMany {
			var HasManyModel *HasManyThingGormModel
			if HasManyModel, err = item.ToModel(); err != nil {
				return
			} else {
				theModel.HasMany = append(theModel.HasMany, HasManyModel)
			}
		}
	}

	if len(p.ManyToMany) > 0 {
		theModel.ManyToMany = []*ManyToManyThingGormModel{}
		for _, item := range p.ManyToMany {
			var ManyToManyModel *ManyToManyThingGormModel
			if ManyToManyModel, err = item.ToModel(); err != nil {
				return
			} else {
				theModel.ManyToMany = append(theModel.ManyToMany, ManyToManyModel)
			}
		}
	}

	theModel.IntEnum = int(p.IntEnum)

	theModel.StringEnum = p.StringEnum.String()

	if len(p.IntEnumList) > 0 {
		theModel.IntEnumList = pq.Int32Array{}
		for _, val := range p.IntEnumList {
			theModel.IntEnumList = append(theModel.IntEnumList, int32(val))
		}
	}

	if len(p.StringEnumList) > 0 {
		theModel.StringEnumList = pq.StringArray{}
		for _, val := range p.StringEnumList {
			theModel.StringEnumList = append(theModel.StringEnumList, val.String())
		}
	}

	return
}

type BelongsToThingGormModels []*BelongsToThingGormModel
type BelongsToThingProtos []*BelongsToThing

func (m BelongsToThingGormModels) ToProtos() (protos BelongsToThingProtos, err error) {
	protos = BelongsToThingProtos{}
	for _, model := range m {
		var proto *BelongsToThing
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p BelongsToThingProtos) ToModels() (models BelongsToThingGormModels, err error) {
	models = BelongsToThingGormModels{}
	for _, proto := range p {
		var model *BelongsToThingGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

type BelongsToThingGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`
}

func (m *BelongsToThingGormModel) TableName() string {
	return "belongs_to_things"
}

func (m *BelongsToThingGormModel) ToProto() (theProto *BelongsToThing, err error) {
	if m == nil {
		return
	}
	theProto = &BelongsToThing{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	return
}

func (p *BelongsToThing) ToModel() (theModel *BelongsToThingGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &BelongsToThingGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	return
}

type HasOneThingGormModels []*HasOneThingGormModel
type HasOneThingProtos []*HasOneThing

func (m HasOneThingGormModels) ToProtos() (protos HasOneThingProtos, err error) {
	protos = HasOneThingProtos{}
	for _, model := range m {
		var proto *HasOneThing
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p HasOneThingProtos) ToModels() (models HasOneThingGormModels, err error) {
	models = HasOneThingGormModels{}
	for _, proto := range p {
		var model *HasOneThingGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

type HasOneThingGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`

	// @gotags: fake:"skip"
	ThingId *string `json:"thingId" fake:"skip"`
}

func (m *HasOneThingGormModel) TableName() string {
	return "has_one_things"
}

func (m *HasOneThingGormModel) ToProto() (theProto *HasOneThing, err error) {
	if m == nil {
		return
	}
	theProto = &HasOneThing{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	theProto.ThingId = m.ThingId

	return
}

func (p *HasOneThing) ToModel() (theModel *HasOneThingGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &HasOneThingGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	theModel.ThingId = p.ThingId

	return
}

type HasManyThingGormModels []*HasManyThingGormModel
type HasManyThingProtos []*HasManyThing

func (m HasManyThingGormModels) ToProtos() (protos HasManyThingProtos, err error) {
	protos = HasManyThingProtos{}
	for _, model := range m {
		var proto *HasManyThing
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p HasManyThingProtos) ToModels() (models HasManyThingGormModels, err error) {
	models = HasManyThingGormModels{}
	for _, proto := range p {
		var model *HasManyThingGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

type HasManyThingGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`

	// @gotags: fake:"skip"
	ThingId *string `json:"thingId" fake:"skip"`
}

func (m *HasManyThingGormModel) TableName() string {
	return "has_many_things"
}

func (m *HasManyThingGormModel) ToProto() (theProto *HasManyThing, err error) {
	if m == nil {
		return
	}
	theProto = &HasManyThing{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	theProto.ThingId = m.ThingId

	return
}

func (p *HasManyThing) ToModel() (theModel *HasManyThingGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &HasManyThingGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	theModel.ThingId = p.ThingId

	return
}

type ManyToManyThingGormModels []*ManyToManyThingGormModel
type ManyToManyThingProtos []*ManyToManyThing

func (m ManyToManyThingGormModels) ToProtos() (protos ManyToManyThingProtos, err error) {
	protos = ManyToManyThingProtos{}
	for _, model := range m {
		var proto *ManyToManyThing
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p ManyToManyThingProtos) ToModels() (models ManyToManyThingGormModels, err error) {
	models = ManyToManyThingGormModels{}
	for _, proto := range p {
		var model *ManyToManyThingGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

type ManyToManyThingGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"createdAt" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;default:now();" json:"updatedAt" fake:"skip"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`
}

func (m *ManyToManyThingGormModel) TableName() string {
	return "many_to_many_things"
}

func (m *ManyToManyThingGormModel) ToProto() (theProto *ManyToManyThing, err error) {
	if m == nil {
		return
	}
	theProto = &ManyToManyThing{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.Name = m.Name

	return
}

func (p *ManyToManyThing) ToModel() (theModel *ManyToManyThingGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &ManyToManyThingGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.Name = p.Name

	return
}
