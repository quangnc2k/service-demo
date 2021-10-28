package services

import (
	"context"
	"git.cyradar.com/phinc/my-awesome-project/internal/model"
	"git.cyradar.com/phinc/my-awesome-project/internal/persistance"
)

func AddStaff(ctx context.Context, staff *model.Staff) error {
	if err := persistance.DefaultRepository.StaffRepository.Insert(ctx, staff); err != nil {
		return err
	}

	return nil
}

func UpdateStaff(ctx context.Context, staff *model.Staff, id string) error {
	current, err := persistance.DefaultRepository.StaffRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	current.Name = staff.Name
	current.DateOfBirth = staff.DateOfBirth
	current.Address = staff.Address

	if err = persistance.DefaultRepository.StaffRepository.Update(ctx, id, &current); err != nil {
		return err
	}

	return nil
}

func RemoveStaff(ctx context.Context, id string) error {
	_, err := persistance.DefaultRepository.StaffRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if err = persistance.DefaultRepository.StaffRepository.Remove(ctx, id); err != nil {
		return err
	}

	err = DefaultDispatcher.DispatchDeleteStaff(id)
	if err != nil {
		return err
	}

	return nil
}

func ViewStaff(ctx context.Context, id string) (*model.Staff, error) {
	found, err := persistance.DefaultRepository.StaffRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &found, nil
}