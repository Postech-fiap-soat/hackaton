package point_record

import "hackaton/internal/app/domain"

type RegisterPointUseCase struct {
	recordPointRepository domain.PointRecordRepository
}

func NewRegisterPointUseCase(recordPointRepository domain.PointRecordRepository) *RegisterPointUseCase {
	return &RegisterPointUseCase{recordPointRepository: recordPointRepository}
}

func (r *RegisterPointUseCase) Handle(registerPointDTO domain.RegisterPointDTO) (*domain.PointRecord, error) {
	point, err := r.recordPointRepository.RegisterPoint(&domain.PointRecord{})
	if err != nil {
		return nil, err
	}
	return point, nil
}
