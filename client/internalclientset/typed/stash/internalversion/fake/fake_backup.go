/*
Copyright 2018 The Stash Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	stash "github.com/appscode/stash/apis/stash"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackups implements BackupInterface
type FakeBackups struct {
	Fake *FakeStash
	ns   string
}

var backupsResource = schema.GroupVersionResource{Group: "stash.appscode.com", Version: "", Resource: "backups"}

var backupsKind = schema.GroupVersionKind{Group: "stash.appscode.com", Version: "", Kind: "Backup"}

// Get takes name of the backup, and returns the corresponding backup object, and an error if there is any.
func (c *FakeBackups) Get(name string, options v1.GetOptions) (result *stash.Backup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupsResource, c.ns, name), &stash.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stash.Backup), err
}

// List takes label and field selectors, and returns the list of Backups that match those selectors.
func (c *FakeBackups) List(opts v1.ListOptions) (result *stash.BackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupsResource, backupsKind, c.ns, opts), &stash.BackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &stash.BackupList{}
	for _, item := range obj.(*stash.BackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backups.
func (c *FakeBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupsResource, c.ns, opts))

}

// Create takes the representation of a backup and creates it.  Returns the server's representation of the backup, and an error, if there is any.
func (c *FakeBackups) Create(backup *stash.Backup) (result *stash.Backup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupsResource, c.ns, backup), &stash.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stash.Backup), err
}

// Update takes the representation of a backup and updates it. Returns the server's representation of the backup, and an error, if there is any.
func (c *FakeBackups) Update(backup *stash.Backup) (result *stash.Backup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupsResource, c.ns, backup), &stash.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stash.Backup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackups) UpdateStatus(backup *stash.Backup) (*stash.Backup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupsResource, "status", c.ns, backup), &stash.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stash.Backup), err
}

// Delete takes name of the backup and deletes it. Returns an error if one occurs.
func (c *FakeBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupsResource, c.ns, name), &stash.Backup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &stash.BackupList{})
	return err
}

// Patch applies the patch and returns the patched backup.
func (c *FakeBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *stash.Backup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupsResource, c.ns, name, data, subresources...), &stash.Backup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*stash.Backup), err
}